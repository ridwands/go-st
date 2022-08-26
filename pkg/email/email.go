package email

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gopkg.in/gomail.v2"
)

func SendEmail(payload map[string]string) {
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", viper.GetString("SENDER_EMAIL"))
	mailer.SetHeader("To", payload["to"])
	mailer.SetHeader("Subject", payload["subject"])

	mailer.SetBody("text/html", payload["body"])

	dialer := gomail.NewDialer(
		viper.GetString("SMTP_HOST"),
		viper.GetInt("SMTP_PORT"),
		viper.GetString("SMTP_AUTH_EMAIL"),
		viper.GetString("SMTP_AUTH_PASSWORD"),
	)

	err := dialer.DialAndSend(mailer)
	if err != nil {
		logrus.Error(err.Error())
	}

	logrus.Info("Mail Sent")
}
