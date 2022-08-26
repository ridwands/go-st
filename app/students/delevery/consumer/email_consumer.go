package consumer

import (
	"encoding/json"
	"github.com/nsqio/go-nsq"
	"github.com/sirupsen/logrus"
	nsqApp "golang/app/nsq"
	"golang/domain"
)

type EmailConsumer struct {
	EmailConsumerUseCase domain.EmailConsumerUsecase
	NSQConsumer          *nsq.Consumer
}

func NewEmailConsumer(UseCase domain.EmailConsumerUsecase, consumer *nsq.Consumer) EmailConsumer {
	return EmailConsumer{
		EmailConsumerUseCase: UseCase,
		NSQConsumer:          consumer,
	}
}

func (h EmailConsumer) StartEmailConsumer(u domain.EmailConsumerUsecase) {
	h.NSQConsumer.AddHandler(&EmailConsumer{
		EmailConsumerUseCase: u,
	})

	nsqApp.NSQConsumerConnection(h.NSQConsumer)
}

func (h EmailConsumer) HandleMessage(m *nsq.Message) error {
	logrus.Info("Message Incoming")
	var payload domain.EmailConsumer
	err := json.Unmarshal(m.Body, &payload)
	if err != nil {
		logrus.Error(err.Error())
		return nil
	}

	_ = h.EmailConsumerUseCase.SendEmailRegistration(payload)
	return nil
}
