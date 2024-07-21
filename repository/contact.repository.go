package repository

import (
	"btb-service/model"
	"btb-service/pkg"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var mongodbContactRepository = pkg.MongoDBDatabase{DatabaseName: "btb_client", CollectionName: "contacts"}

func GetContacts() (data []map[string]interface{}, err error) {
	mongodbContactRepository.Filter = bson.D{{}}
	data, err = mongodbContactRepository.GetMongoDB()
	return
}

func SaveContacts(payload model.ContactInsertPayload) (err error) {
	payload.RegisteredDate = primitive.NewDateTimeFromTime(time.Now())
	data, _ := pkg.StructToMap(payload)
	mongodbContactRepository.Payload = data
	err = mongodbContactRepository.InsertMongoDB()
	return
}
