package pkg

import (
	"errors"
	"strconv"

	"go.mongodb.org/mongo-driver/bson"
	"gopkg.in/gomail.v2"
)

type MailPayload struct {
	To      []string
	Cc      []string
	Subject string
	Message string
}

func (payload MailPayload) SendMail() (err error) {
	var data map[string]interface{}
	var mongoDBSecretKeyRepository = MongoDBDatabase{DatabaseName: "btb_app", CollectionName: "secretkey"}
	mongoDBSecretKeyRepository.Filter = bson.M{"type": "emailconfig"}
	queryData, err := mongoDBSecretKeyRepository.GetMongoDB()
	if err != nil {
		return
	}

	if len(queryData) == 0 {
		err = errors.New("email config not found")
		return
	}
	data = queryData[0]

	smtpemail, ok := data["smtpemail"].(string)
	if !ok {
		err = errors.New("failed to get smtpemail")
		return
	}

	smtppassword, ok := data["smtppassword"].(string)
	if !ok {
		err = errors.New("failed to get smtppassword")
		return
	}

	smtphost, ok := data["smtphost"].(string)
	if !ok {
		err = errors.New("failed to get smtphost")
		return
	}

	smtpport, ok := data["smtpport"].(string)
	if !ok {
		err = errors.New("failed to get smtpport")
		return
	}

	smtpPort, err := strconv.Atoi(smtpport)
	if err != nil {
		return
	}

	m := gomail.NewMessage()
	m.SetHeader("From", smtpemail)
	if len(payload.To) == 0 {
		err = errors.New("destination is empty")
		return
	} else {
		m.SetHeader("To", payload.To...)
	}
	if len(payload.Cc) != 0 {
		m.SetHeader("Cc", payload.Cc...)
	}
	m.SetHeader("Subject", payload.Subject)
	m.SetBody("text/html", payload.Message)

	d := gomail.NewDialer(smtphost, smtpPort, smtpemail, smtppassword)
	return d.DialAndSend(m)
}
