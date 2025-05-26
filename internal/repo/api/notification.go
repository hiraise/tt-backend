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
	confirmationUrl string
}

func NewSmtpNotificationRepo(
	sender smtp.Sender,
	logger logger.Logger,
	uuidGenerator uuid.Generator,
	confirmationUrl string,
) *SmtpNotificationRepo {
	return &SmtpNotificationRepo{
		sender:          sender,
		logger:          logger,
		uuidGenerator:   uuidGenerator,
		confirmationUrl: confirmationUrl,
	}
}

func (r *SmtpNotificationRepo) SendConfirmationEmail(ctx context.Context, email string, token string) error {
	msg := smtp.Message{
		Recipients: []string{email},
		Subject:    "Account confirmation",
		Text:       r.confirmationUrl + token,
	}
	eventId := r.uuidGenerator.Generate()
	if err := r.sender.Send(msg, eventId); err != nil {
		return repo.Wrap(repo.ErrInternal, err)
	}
	return nil
}
