package repository

import (
	"btb-service/model"
	"btb-service/pkg"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var mongodbCareerRepository = pkg.MongoDBDatabase{DatabaseName: "btb_client"}

func GetCareerApplicantData() (data []map[string]interface{}, err error) {
	mongodbCareerRepository.CollectionName = "career_applicant"
	mongodbCareerRepository.Filter = bson.D{{}}
	mongodbCareerRepository.Sort = map[string]interface{}{"registereddate": -1}
	data, err = mongodbCareerRepository.GetMongoDB()
	return
}

func ApplyCareer(payload model.CareerApplyInsertPayload) (err error) {
	mongodbCareerRepository.CollectionName = "career_applicant"
	data, _ := pkg.StructToMap(payload)
	data["registereddate"] = primitive.NewDateTimeFromTime(time.Now())
	mongodbCareerRepository.Payload = data
	err = mongodbCareerRepository.InsertMongoDB()
	return
}

func UpsertCareer(payload model.CareerUpsertPayload) (err error) {
	mongodbCareerRepository.CollectionName = "career"
	if pkg.IsEmptyString(payload.ID) {
		payload.RegisteredDate = primitive.NewDateTimeFromTime(time.Now())
		data, _ := pkg.StructToMap(payload)
		mongodbCareerRepository.Payload = data
		err = mongodbCareerRepository.InsertMongoDB()
	} else {
		data, _ := pkg.StructToMap(payload)
		mongodbCareerRepository.Payload = data
		delete(mongodbCareerRepository.Payload, "_id")
		id, _ := primitive.ObjectIDFromHex(payload.ID)
		mongodbCareerRepository.Filter = bson.D{{Key: "_id", Value: id}}
		err = mongodbCareerRepository.UpdateMongoDB()
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

	mongodbCareerRepository.Sort = map[string]interface{}{"registereddate": -1}
	queryData, err := mongodbCareerRepository.GetMongoDB()
	if err != nil {
		log.Println(err.Error())
		return
	}

	for _, v := range queryData {
		maximumApplyDate, ok := v["maximumApplyDate"].(string)
		if ok {
			compareDate, err := pkg.CompareIsoDateStringToNow(maximumApplyDate)
			if err != nil {
				log.Println(err.Error())
				continue
			}

			if compareDate >= 0 {
				data = append(data, v)
			}
		} else {
			log.Println("not ok")
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
	mongodbCareerRepository.Sort = map[string]interface{}{"registereddate": -1}
	data, err = mongodbCareerRepository.GetMongoDB()
	return
}

func DeleteCareer(data model.CareerDeletePayload) (err error) {
	mongodbCareerRepository.CollectionName = "career"
	id, _ := primitive.ObjectIDFromHex(data.ID)
	delete(mongodbCareerRepository.Payload, "_id")
	mongodbCareerRepository.Filter = bson.D{{Key: "_id", Value: id}}
	err = mongodbCareerRepository.DeleteMongoDB()
	return
}
