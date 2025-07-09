package project

import (
	"context"
	"errors"
	"fmt"
	"slices"
	"task-trail/internal/repo"
	"task-trail/internal/usecase/dto"
)

func (u *UseCase) AddMembers(ctx context.Context, data *dto.ProjectAddMembers) error {
	project, err := u.getOwnedProject(ctx, data.ProjectID, data.OwnerID)
	if err != nil {
		return err
	}

	if err := u.verifyNewMembers(project.Members, data.MemberEmails); err != nil {
		return err
	}
	newEmails, err := u.getUnregisteredEmails(ctx, data.MemberEmails)

	if err != nil {
		return u.errHandler.InternalTrouble(err, "project members loading failed", "projectID", data.ProjectID)
	}
	fmt.Println(newEmails, project)
	return nil
}

func (u *UseCase) getOwnedProject(ctx context.Context, projectID int, ownerID int) (*dto.Project, error) {
	project, err := u.projectRepo.GetOwnedProject(ctx, projectID, ownerID)
	if err != nil {
		if errors.Is(err, repo.ErrNotFound) {
			return nil, u.errHandler.NotFound(err, "project not found", "projectID", projectID, "ownerID", ownerID)
		}
		return nil, u.errHandler.InternalTrouble(err, "project loading failed", "projectID", projectID, "ownerID", ownerID)
	}
	return project, nil
}

func (u *UseCase) verifyNewMembers(pMembers []*dto.UserEmailAndID, newMembers []string) error {
	for _, v := range pMembers {
		if slices.Contains(newMembers, v.Email) {
			return u.errHandler.BadRequest(nil, "member already in project", "memberEmail", v.Email)
		}
	}
	return nil
}

func (u *UseCase) getUnregisteredEmails(ctx context.Context, newMembers []string) ([]string, error) {
	f, err := u.userRepo.GetIdsByEmails(ctx, newMembers)
	if err != nil {
		return nil, err
	}
	var foundEmails = make(map[string]struct{}, len(f))
	for _, v := range f {
		foundEmails[v.Email] = struct{}{}
	}
	var newUsers []string
	for _, email := range newMembers {
		if _, ok := foundEmails[email]; !ok {
			newUsers = append(newUsers, email)
		}
	}
	return newUsers, nil
}
