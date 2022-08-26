package domain

type EmailConsumer struct {
	Name  string `json:"name"`
	NISN  string `json:"nisn"`
	Email string `json:"email"`
}

type EmailConsumerUsecase interface {
	SendEmailRegistration(payload EmailConsumer) interface{}
}
