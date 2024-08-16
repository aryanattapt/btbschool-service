package pkg

import (
	"errors"
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
	if len(payload.To) == 0 {
		err = errors.New("destination is empty")
		return
	} else {
		m.SetHeader("To", strings.Join(payload.To, ","))
	}
	if len(payload.Cc) != 0 {
		m.SetHeader("Cc", strings.Join(payload.Cc, ","))
	}
	m.SetHeader("Subject", payload.Subject)
	m.SetBody("text/html", payload.Message)

	smtpPort, err := strconv.Atoi(os.Getenv("MAIL_SMTP_PORT"))
	if err != nil {
		return
	}

	d := gomail.NewDialer(os.Getenv("MAIL_SMTP_HOST"), smtpPort, os.Getenv("MAIL_AUTH_EMAIL"), os.Getenv("MAIL_AUTH_PASSWORD"))
	return d.DialAndSend(m)
}
