package pkg

import (
	"fmt"
	"net/smtp"
	"os"
	"strings"
)

var (
	MAIL_SMTP_HOST     = os.Getenv("MAIL_SMTP_HOST")
	MAIL_SMTP_PORT     = os.Getenv("MAIL_SMTP_PORT")
	MAIL_SENDER_NAME   = os.Getenv("MAIL_SENDER_NAME")
	MAIL_AUTH_EMAIL    = os.Getenv("MAIL_AUTH_EMAIL")
	MAIL_AUTH_PASSWORD = os.Getenv("MAIL_AUTH_PASSWORD")
)

type MailPayload struct {
	To      []string
	Cc      []string
	Subject string
	Message string
}

func (payload MailPayload) SendMail() (err error) {
	body := "From: " + MAIL_SENDER_NAME + "\n" +
		"To: " + strings.Join(payload.To, ",") + "\n" +
		"Cc: " + strings.Join(payload.Cc, ",") + "\n" +
		"Subject: " + payload.Subject + "\n\n" +
		payload.Message

	auth := smtp.PlainAuth("", MAIL_AUTH_EMAIL, MAIL_AUTH_PASSWORD, MAIL_SMTP_HOST)
	smtpAddr := fmt.Sprintf("%s:%s", MAIL_SMTP_HOST, MAIL_SMTP_PORT)

	err = smtp.SendMail(smtpAddr, auth, MAIL_AUTH_EMAIL, append(payload.To, payload.Cc...), []byte(body))
	if err != nil {
		return err
	}

	return nil
}
