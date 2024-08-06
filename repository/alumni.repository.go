package repository

import (
	"btb-service/model"
	"btb-service/pkg"

	"go.mongodb.org/mongo-driver/bson"
)

var mongodbAlumniRepository = pkg.MongoDBDatabase{DatabaseName: "btb_client", CollectionName: "alumni"}

func GetAlumni() (data []map[string]interface{}, err error) {
	mongodbAlumniRepository.Filter = bson.D{{}}
	data, err = mongodbAlumniRepository.GetMongoDB()
	return
}

func SaveAlumni(payload model.AlumniInsertPayload) (err error) {
	data, _ := pkg.StructToMap(payload)
	mongodbAlumniRepository.Payload = data
	err = mongodbAlumniRepository.InsertMongoDB()
	return
}
