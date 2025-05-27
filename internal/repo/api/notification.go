package api

import (
	"context"
	"task-trail/internal/pkg/logger"
	"task-trail/internal/pkg/smtp"
	"task-trail/internal/pkg/uuid"
	"task-trail/internal/repo"
)

type SmtpNotificationRepo struct {
	sender          smtp.Sender
	logger          logger.Logger
	uuidGenerator   uuid.Generator
	verificationUrl string
}

func NewSmtpNotificationRepo(
	sender smtp.Sender,
	logger logger.Logger,
	uuidGenerator uuid.Generator,
	verificationUrl string,
) *SmtpNotificationRepo {
	return &SmtpNotificationRepo{
		sender:          sender,
		logger:          logger,
		uuidGenerator:   uuidGenerator,
		verificationUrl: verificationUrl,
	}
}

func (r *SmtpNotificationRepo) SendVerificationEmail(ctx context.Context, email string, token string) error {
	msg := smtp.Message{
		Recipients: []string{email},
		Subject:    "Account Verification",
		Text:       r.verificationUrl + token,
	}
	eventID := r.uuidGenerator.Generate()
	if err := r.sender.Send(msg, eventID); err != nil {
		return repo.Wrap(repo.ErrInternal, err)
	}
	return nil
}
