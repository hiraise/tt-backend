//go:build integration

package persistent

import (
	"context"
	"task-trail/internal/entity"
	"task-trail/internal/repo"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func beforeEachEmailTest(t *testing.T) {
	cleanDB(t)
	err := userRepo.Create(t.Context(), &basicUser)
	require.NoError(t, err)
}

func createEmailConfirmToken(ctx context.Context, userId int, tokenId string) error {
	e := entity.EmailToken{
		ID:        tokenId,
		ExpiredAt: time.Now().Add(time.Minute * 10),
		UserId:    userId,
		Purpose:   entity.PurposeConfirmation,
	}
	return emailTokenRepo.Create(ctx, e)
}
func TestCreateConfirmationToken(t *testing.T) {
	ctx := t.Context()
	beforeEachEmailTest(t)

	t.Run("successfully create", func(t *testing.T) {
		err := createEmailConfirmToken(ctx, 1, testTokenId)
		require.NoError(t, err)
	})
	t.Run("already exists", func(t *testing.T) {
		err := createEmailConfirmToken(ctx, 1, testTokenId)
		require.ErrorIs(t, err, repo.ErrConflict)
	})
	t.Run("user not found", func(t *testing.T) {
		err := createEmailConfirmToken(ctx, 2, testTokenId1)
		require.ErrorIs(t, err, repo.ErrNotFound)
	})
	t.Run("database internal error", func(t *testing.T) {
		err := createEmailConfirmToken(getBadContext(t), 1, testTokenId)
		require.ErrorIs(t, err, repo.ErrInternal)
	})

}

func TestEmailTokenGetById(t *testing.T) {
	ctx := t.Context()
	beforeEachEmailTest(t)
	err := createEmailConfirmToken(ctx, 1, testTokenId)
	require.NoError(t, err)

	t.Run("successfully get confirmation token by ID", func(t *testing.T) {
		token, err := emailTokenRepo.GetById(ctx, testTokenId)
		require.NoError(t, err)
		require.Equal(t, testTokenId, token.ID)
		require.Equal(t, entity.PurposeConfirmation, token.Purpose)
	})
	t.Run("token not found", func(t *testing.T) {
		token, err := emailTokenRepo.GetById(ctx, testTokenId1)
		require.Nil(t, token)
		require.ErrorIs(t, err, repo.ErrNotFound)
	})
	t.Run("database internal error", func(t *testing.T) {
		token, err := emailTokenRepo.GetById(getBadContext(t), testTokenId)
		require.Nil(t, token)
		require.ErrorIs(t, err, repo.ErrInternal)
	})
}

func TestEmailTokenUse(t *testing.T) {
	ctx := t.Context()
	beforeEachEmailTest(t)
	err := createEmailConfirmToken(ctx, 1, testTokenId)
	require.NoError(t, err)
	t.Run("successfully use token", func(t *testing.T) {
		err := emailTokenRepo.Use(ctx, testTokenId)
		require.NoError(t, err)
		token, err := emailTokenRepo.GetById(ctx, testTokenId)
		require.NoError(t, err)
		require.NotNil(t, token.UsedAt)
	})
	t.Run("token not found", func(t *testing.T) {
		err := emailTokenRepo.Use(ctx, testTokenId1)
		require.ErrorIs(t, err, repo.ErrNotFound)
	})
	t.Run("token already used", func(t *testing.T) {
		err := emailTokenRepo.Use(ctx, testTokenId)
		require.ErrorIs(t, err, repo.ErrNotFound)
	})
	t.Run("database internal error", func(t *testing.T) {
		err := emailTokenRepo.Use(getBadContext(t), testTokenId)
		require.ErrorIs(t, err, repo.ErrInternal)
	})
}
