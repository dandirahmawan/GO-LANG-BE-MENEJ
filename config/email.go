package config

import (
	"log"

	"gopkg.in/gomail.v2"
)

// const CONFIG_SMTP_HOST = "smtp.gmail.com"
// const CONFIG_SMTP_PORT = 587
// const CONFIG_SENDER_NAME = "Menej"
// const CONFIG_AUTH_EMAIL = "menejproyek@gmail.com"
// const CONFIG_AUTH_PASSWORD = "majulancar1"

const CONFIG_SMTP_HOST = "smtp-mail.outlook.com"
const CONFIG_SMTP_PORT = 587
const CONFIG_SENDER_NAME = "Menej"
const CONFIG_AUTH_PASSWORD = "majulancar1"
const CONFIG_AUTH_EMAIL = "dandirahmawan95@outlook.com"

func SendMail(to []string, cc []string, subject, message string) error {
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", CONFIG_AUTH_EMAIL)
	mailer.SetHeader("To", to...)
	mailer.SetHeader("Subject", subject)
	mailer.SetBody("text/html", message)

	dialer := gomail.NewDialer(
		CONFIG_SMTP_HOST,
		CONFIG_SMTP_PORT,
		CONFIG_AUTH_EMAIL,
		CONFIG_AUTH_PASSWORD,
	)

	err := dialer.DialAndSend(mailer)
	if err != nil {
		log.Fatal(err.Error())
		return err
	}

	return nil
}
