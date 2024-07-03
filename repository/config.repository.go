package repository

import (
	"btb-service/model"
	"btb-service/pkg"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var mongoDBParam = pkg.MongoDBDatabase{DatabaseName: "client_configs"}

type ConfigRepositoryModel struct {
	model.ConfigModel
}

func (payload ConfigRepositoryModel) GetConfig() (data []map[string]interface{}, err error) {
	mongoDBParam.CollectionName = payload.Type
	mongoDBParam.Filter = bson.D{{}}
	data, err = mongoDBParam.GetMongoDB()
	return
}

func (payload ConfigRepositoryModel) UpsertConfig() (err error) {
	mongoDBParam.CollectionName = payload.Type
	for _, payload := range payload.Payload {
		idPayload, ok := payload["_id"].(string)
		mongoDBParam.Payload = payload

		if pkg.IsEmptyString(idPayload) || !ok {
			err = mongoDBParam.InsertMongoDB()
		} else {
			delete(mongoDBParam.Payload, "_id")
			id, _ := primitive.ObjectIDFromHex(idPayload)
			mongoDBParam.Filter = bson.D{{Key: "_id", Value: id}}
			err = mongoDBParam.UpdateMongoDB()
		}

		if err != nil {
			return
		}
	}
	return
}

func (payload ConfigRepositoryModel) DeleteConfig() (err error) {
	mongoDBParam.CollectionName = payload.Type
	for _, payload := range payload.Payload {
		var idPayload string = payload["_id"].(string)
		mongoDBParam.Payload = payload

		if pkg.IsEmptyString(idPayload) {
			err = errors.New("ID is mandatory")
		} else {
			delete(mongoDBParam.Payload, "_id")
			id, _ := primitive.ObjectIDFromHex(idPayload)
			mongoDBParam.Filter = bson.D{{Key: "_id", Value: id}}
			err = mongoDBParam.DeleteMongoDB()
		}

		if err != nil {
			return
		}
	}
	return
}
