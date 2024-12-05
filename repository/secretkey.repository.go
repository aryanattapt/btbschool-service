package repository

import (
	"btb-service/model"
	"btb-service/pkg"
	"errors"

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
