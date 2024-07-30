package repository

import (
	"btb-service/model"
	"btb-service/pkg"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var mongodbCareerRepository = pkg.MongoDBDatabase{DatabaseName: "btb_client"}

func GetCareerApplicantData() (data []map[string]interface{}, err error) {
	mongodbCareerRepository.CollectionName = "career_applicant"
	mongodbCareerRepository.Filter = bson.D{{}}
	data, err = mongodbCareerRepository.GetMongoDB()
	return
}

func ApplyCareer(payload model.CareerApplyInsertPayload) (err error) {
	mongodbCareerRepository.CollectionName = "career_applicant"
	data, _ := pkg.StructToMap(payload)
	mongodbCareerRepository.Payload = data
	err = mongodbCareerRepository.InsertMongoDB()
	return
}

func UpsertCareer(payload model.CareerUpsertPayload) (err error) {
	mongodbCareerRepository.CollectionName = "career"
	data, _ := pkg.StructToMap(payload)
	mongoDBConfigRepository.Payload = data
	if pkg.IsEmptyString(payload.ID) {
		err = mongoDBConfigRepository.InsertMongoDB()
	} else {
		delete(mongoDBConfigRepository.Payload, "_id")
		id, _ := primitive.ObjectIDFromHex(payload.ID)
		mongoDBConfigRepository.Filter = bson.D{{Key: "_id", Value: id}}
		err = mongoDBConfigRepository.UpdateMongoDB()
	}
	return
}

func GetActiveCareer(searchPayload map[string]interface{}) (data []map[string]interface{}, err error) {
	mongodbCareerRepository.CollectionName = "career"
	idPayload, ok := searchPayload["_id"].(string)
	if ok {
		id, _ := primitive.ObjectIDFromHex(idPayload)
		delete(searchPayload, "_id")
		mongodbCareerRepository.Filter = bson.D{{Key: "_id", Value: id}}
	} else {
		mongodbCareerRepository.Filter = searchPayload
	}

	queryData, err := mongodbCareerRepository.GetMongoDB()
	if err != nil {
		return
	}

	for _, v := range queryData {
		maximumApplyDate := v["maximumApplyDate"].(string)
		compareDate, err := pkg.CompareIsoDateStringToNow(maximumApplyDate)
		if err != nil && compareDate >= 0 {
			data = append(data, v)
		} else {
			log.Println("hasil compare date: ", compareDate)
			log.Println(err.Error())
		}
	}

	return
}

func GetAllCareer(searchPayload map[string]interface{}) (data []map[string]interface{}, err error) {
	mongodbCareerRepository.CollectionName = "career"
	idPayload, ok := searchPayload["_id"].(string)
	if ok {
		id, _ := primitive.ObjectIDFromHex(idPayload)
		delete(searchPayload, "_id")
		mongodbCareerRepository.Filter = bson.D{{Key: "_id", Value: id}}
	} else {
		mongodbCareerRepository.Filter = searchPayload
	}
	data, err = mongodbCareerRepository.GetMongoDB()
	return
}

func DeleteCareer(data model.CareerDeletePayload) (err error) {
	mongoDBConfigRepository.CollectionName = "career"
	id, _ := primitive.ObjectIDFromHex(data.ID)
	delete(mongoDBConfigRepository.Payload, "_id")
	mongoDBConfigRepository.Filter = bson.D{{Key: "_id", Value: id}}
	err = mongoDBConfigRepository.DeleteMongoDB()
	return
}
