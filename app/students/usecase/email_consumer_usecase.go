package usecase

import (
	"golang/domain"
	"golang/pkg/email"
)

type EmailConsumerUsecase struct {
}

func NewEmailConsumerUsecase() domain.EmailConsumerUsecase {
	return &EmailConsumerUsecase{}
}

func (e EmailConsumerUsecase) SendEmailRegistration(payload domain.EmailConsumer) interface{} {
	m := map[string]string{
		"to":      payload.Email,
		"subject": "Students Registration",
		"body":    "Hello " + payload.Name + " <b>have a nice day</b>",
	}
	email.SendEmail(m)
	return nil
}
