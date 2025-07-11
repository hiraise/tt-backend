package project

import (
	"context"
	"errors"
	"task-trail/internal/repo"
	"task-trail/internal/usecase/dto"
)

func (u *UseCase) GetCandidates(ctx context.Context, ownerID int, projectID int) ([]*dto.UserSimple, error) {
	if projectID != 0 {
		if err := u.projectRepo.IsMember(ctx, projectID, ownerID); err != nil {
			if errors.Is(err, repo.ErrNotFound) {
				return nil, u.errHandler.NotFound(
					err,
					"project not found",
					"ownerID", ownerID,
					"projectID", projectID,
				)
			}
			return nil, u.errHandler.InternalTrouble(
				err,
				"failed to verify user membership",
				"ownerID", ownerID,
				"projectID", projectID,
			)
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
