package repository

import (
	"btb-service/model"
	"btb-service/pkg"
	"errors"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var mongoDBConfigRepository = pkg.MongoDBDatabase{DatabaseName: "btb_app"}

type ConfigRepositoryModel struct {
	model.ConfigModel
}

func (payload ConfigRepositoryModel) GetConfig(searchPayload map[string]interface{}) (data []map[string]interface{}, err error) {
	idPayload, ok := searchPayload["_id"].(string)
	if ok {
		id, _ := primitive.ObjectIDFromHex(idPayload)
		delete(searchPayload, "_id")
		mongoDBConfigRepository.Filter = bson.D{{Key: "_id", Value: id}}
	} else {
		mongoDBConfigRepository.Filter = searchPayload
	}

	mongoDBConfigRepository.CollectionName = payload.Type
	data, err = mongoDBConfigRepository.GetMongoDB()
	return
}

func (payload ConfigRepositoryModel) UpsertConfig() (err error) {
	mongoDBConfigRepository.CollectionName = payload.Type
	for _, payload := range payload.Payload {
		idPayload, ok := payload["_id"].(string)
		mongoDBConfigRepository.Payload = payload
		log.Println(idPayload)

		if pkg.IsEmptyString(idPayload) || !ok {
			err = mongoDBConfigRepository.InsertMongoDB()
		} else {
			delete(mongoDBConfigRepository.Payload, "_id")
			id, errID := primitive.ObjectIDFromHex(idPayload)
			if errID != nil {
				return errID
			}
			mongoDBConfigRepository.Filter = bson.D{{Key: "_id", Value: id}}
			err = mongoDBConfigRepository.UpdateMongoDB()
		}

		if err != nil {
			return
		}
	}
	return
}

func (payload ConfigRepositoryModel) DeleteConfig() (err error) {
	mongoDBConfigRepository.CollectionName = payload.Type
	for _, payload := range payload.Payload {
		var idPayload string = payload["_id"].(string)
		mongoDBConfigRepository.Payload = payload

		if pkg.IsEmptyString(idPayload) {
			err = errors.New("ID is mandatory")
		} else {
			delete(mongoDBConfigRepository.Payload, "_id")
			id, _ := primitive.ObjectIDFromHex(idPayload)
			mongoDBConfigRepository.Filter = bson.D{{Key: "_id", Value: id}}
			err = mongoDBConfigRepository.DeleteMongoDB()
		}

		if err != nil {
			return
		}
	}
	return
}
