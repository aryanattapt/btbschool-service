package pkg

import (
	"os"
	"strconv"
	"strings"

	"gopkg.in/gomail.v2"
)

type MailPayload struct {
	To      []string
	Cc      []string
	Subject string
	Message string
}

func (payload MailPayload) SendMail() (err error) {
	m := gomail.NewMessage()
	m.SetHeader("From", os.Getenv("MAIL_AUTH_EMAIL"))
	m.SetHeader("To", strings.Join(payload.To, ","))
	m.SetHeader("Cc", strings.Join(payload.Cc, ","))
	m.SetHeader("Subject", payload.Subject)
	m.SetBody("text/plain", payload.Message)

	smtpPort, err := strconv.Atoi(os.Getenv("MAIL_SMTP_PORT"))
	if err != nil {
		return
	}

	d := gomail.NewDialer(os.Getenv("MAIL_SMTP_HOST"), smtpPort, os.Getenv("MAIL_AUTH_EMAIL"), os.Getenv("MAIL_AUTH_PASSWORD"))
	return d.DialAndSend(m)
}
