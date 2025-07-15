package project

import (
	"context"
	"errors"
	"task-trail/internal/repo"
	"task-trail/internal/usecase/dto"
)

func (u *UseCase) GetCandidates(ctx context.Context, ownerID int, projectID int) ([]*dto.UserSimple, error) {
	// if project id is passed, verify users membership
	if projectID != 0 {
		if err := u.CheckMembership(ctx, projectID, ownerID); err != nil {
			return nil, err
		}
	}
	res, err := u.projectRepo.GetCandidates(ctx, ownerID, projectID)
	if err != nil {
		return nil, u.errHandler.InternalTrouble(
			err,
			"failed to get candidates",
			"ownerID", ownerID,
			"projectID", projectID,
		)
	}
	return res, nil
}

func (u *UseCase) CheckMembership(ctx context.Context, projectID int, memberID int) error {
	if err := u.projectRepo.IsMember(ctx, projectID, memberID); err != nil {
		if errors.Is(err, repo.ErrNotFound) {
			return u.errHandler.NotFound(
				err,
				"project not found",
				"memberID", memberID,
				"projectID", projectID,
			)
		}
		return u.errHandler.InternalTrouble(
			err,
			"failed to verify user membership",
			"memberID", memberID,
			"projectID", projectID,
		)
	}
	return nil
}
