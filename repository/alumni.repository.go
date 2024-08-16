package repository

import (
	"btb-service/model"
	"btb-service/pkg"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var mongodbAlumniRepository = pkg.MongoDBDatabase{DatabaseName: "btb_client", CollectionName: "alumni"}

func GetAlumni(payload map[string]interface{}) (data []map[string]interface{}, err error) {
	idPayload, ok := payload["_id"].(string)
	if ok {
		id, _ := primitive.ObjectIDFromHex(idPayload)
		delete(payload, "_id")
		mongodbAlumniRepository.Filter = bson.D{{Key: "_id", Value: id}}
	} else {
		mongodbAlumniRepository.Filter = payload
	}
	data, err = mongodbAlumniRepository.GetMongoDB()
	return
}

func SaveAlumni(payload model.AlumniInsertPayload) (err error) {
	data, _ := pkg.StructToMap(payload)
	mongodbAlumniRepository.Payload = data
	err = mongodbAlumniRepository.InsertMongoDB()
	return
}

func VerifyAlumni(payload model.AlumniVerifyPayload) (err error) {
	id, _ := primitive.ObjectIDFromHex(payload.ID)
	mongodbAlumniRepository.Filter = bson.D{{Key: "_id", Value: id}}
	data, _ := pkg.StructToMap(payload)
	mongodbAlumniRepository.Payload = data
	delete(mongodbAlumniRepository.Payload, "_id")
	err = mongodbAlumniRepository.UpdateMongoDB()
	return
}
