package repository

import (
	"btb-service/model"
	"btb-service/pkg"
	"errors"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var mongoDBSecretKeyRepository = pkg.MongoDBDatabase{DatabaseName: "btb_app", CollectionName: "secretkey"}

func GetEmailConfig() (data map[string]interface{}, err error) {
	mongoDBSecretKeyRepository.Filter = bson.M{"type": "emailconfig"}
	queryData, err := mongoDBSecretKeyRepository.GetMongoDB()

	if len(queryData) > 0 {
		data = queryData[0]
	} else {
		err = errors.New("email config Not Found")
	}
	return
}

func UpdateEmailConfig(data model.SecretKeyPayload) (err error) {
	idPayload, ok := data.Payload["_id"].(string)
	mongoDBSecretKeyRepository.Payload = data.Payload

	if pkg.IsEmptyString(idPayload) || !ok {
		err = mongoDBSecretKeyRepository.InsertMongoDB()
	} else {
		delete(mongoDBSecretKeyRepository.Payload, "_id")
		id, errID := primitive.ObjectIDFromHex(idPayload)
		if errID != nil {
			return errID
		}
		mongoDBSecretKeyRepository.Filter = bson.M{"type": "emailconfig", "_id": id}
		err = mongoDBSecretKeyRepository.UpdateMongoDB()
	}
	return
}

func GetRecaptchaConfig() (data map[string]interface{}, err error) {
	mongoDBSecretKeyRepository.Filter = bson.M{"type": "recaptcha"}
	queryData, err := mongoDBSecretKeyRepository.GetMongoDB()

	if len(queryData) > 0 {
		data = queryData[0]
	} else {
		err = errors.New("recaptcha config Not Found")
	}
	return
}

func UpdateRecaptchaConfig(data model.SecretKeyPayload) (err error) {
	idPayload, ok := data.Payload["_id"].(string)
	mongoDBSecretKeyRepository.Payload = data.Payload

	if pkg.IsEmptyString(idPayload) || !ok {
		err = mongoDBSecretKeyRepository.InsertMongoDB()
	} else {
		delete(mongoDBSecretKeyRepository.Payload, "_id")
		id, errID := primitive.ObjectIDFromHex(idPayload)
		if errID != nil {
			return errID
		}
		mongoDBSecretKeyRepository.Filter = bson.M{"type": "recaptcha", "_id": id}
		err = mongoDBSecretKeyRepository.UpdateMongoDB()
	}
	return
}

func GetInstagramConfig() (data map[string]interface{}, err error) {
	mongoDBSecretKeyRepository.Filter = bson.M{"type": "instagramtoken"}
	queryData, err := mongoDBSecretKeyRepository.GetMongoDB()

	if len(queryData) > 0 {
		data = queryData[0]
	} else {
		err = errors.New("instragram config Not Found")
	}
	return
}

func UpdateInstagramConfig(data model.SecretKeyPayload) (err error) {
	idPayload, ok := data.Payload["_id"].(string)
	mongoDBSecretKeyRepository.Payload = data.Payload

	if pkg.IsEmptyString(idPayload) || !ok {
		err = mongoDBSecretKeyRepository.InsertMongoDB()
	} else {
		delete(mongoDBSecretKeyRepository.Payload, "_id")
		id, errID := primitive.ObjectIDFromHex(idPayload)
		if errID != nil {
			return errID
		}
		mongoDBSecretKeyRepository.Filter = bson.M{"type": "instagramtoken", "_id": id}
		err = mongoDBSecretKeyRepository.UpdateMongoDB()
	}
	return
}

func GetInstagramToken() (data string, err error) {
	mongoDBSecretKeyRepository.Filter = bson.M{"type": "instagramtoken"}
	queryData, err := mongoDBSecretKeyRepository.GetMongoDB()

	if len(queryData) > 0 {
		data, _ = queryData[0]["token"].(string)
		log.Printf("Token ig: %s", data)
	} else {
		err = errors.New("token Not Found")
	}
	return
}

func UpdateInstagramToken(token string) (err error) {
	log.Println("Update new instagram token : ", token)
	mongoDBSecretKeyRepository.Payload = bson.M{"token": token, "updateddate": pkg.FormatTime(time.Now(), "2006-01-02T15:04:05.999Z")}
	mongoDBSecretKeyRepository.Filter = bson.M{"type": "instagramtoken"}
	err = mongoDBSecretKeyRepository.UpdateMongoDB()
	return
}
