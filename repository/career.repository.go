package repository

import (
	"btb-service/model"
	"btb-service/pkg"

	"go.mongodb.org/mongo-driver/bson"
)

var mongodbCareerRepository = pkg.MongoDBDatabase{DatabaseName: "btb_client", CollectionName: "alumni"}

func GetCareerApplicantData() (data []map[string]interface{}, err error) {
	mongodbCareerRepository.Filter = bson.D{{}}
	data, err = mongodbCareerRepository.GetMongoDB()
	return
}

func ApplyCareer(payload model.CareerApplyInsertPayload) (err error) {
	data, _ := pkg.StructToMap(payload)
	mongodbCareerRepository.Payload = data
	err = mongodbCareerRepository.InsertMongoDB()
	return
}
