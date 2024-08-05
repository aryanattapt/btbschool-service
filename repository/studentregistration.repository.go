package repository

import (
	"btb-service/model"
	"btb-service/pkg"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var mongodbStudentRegistrationRepository = pkg.MongoDBDatabase{DatabaseName: "btb_client", CollectionName: "studentregistration"}

func GetDraftStudentRegistrationData(payload model.DraftStudentRegistrationData) (data []map[string]interface{}, err error) {
	mongodbStudentRegistrationRepository.Filter = bson.M{
		"status": "draft",
		"$and": bson.A{
			payload,
		},
	}
	data, err = mongodbStudentRegistrationRepository.GetMongoDB()
	return
}

func SubmitDataStudentRegistration(payload model.StudentRegistrationInsertPayload) (registrationcode string, err error) {
	if pkg.IsEmptyString(payload.RegistrationCode) {
		registrationcode = "PPD" + time.Now().Format("02012006150405")
		payload.RegisteredDate = primitive.NewDateTimeFromTime(time.Now())
		payload.RegistrationCode = registrationcode
		if pkg.IsEmptyString(payload.Status) {
			payload.Status = "draft"
		}
		data, _ := pkg.StructToMap(payload)
		delete(data, "_id")

		mongodbStudentRegistrationRepository.Payload = data
		err = mongodbStudentRegistrationRepository.InsertMongoDB()
		return
	} else {
		payload.Updateddate = primitive.NewDateTimeFromTime(time.Now())
		registrationcode = payload.RegistrationCode
		mongodbStudentRegistrationRepository.Filter = bson.D{{Key: "registrationcode", Value: payload.RegistrationCode}}
		data, _ := pkg.StructToMap(payload)
		mongodbStudentRegistrationRepository.Payload = data
		err = mongodbStudentRegistrationRepository.UpdateMongoDB()
		return
	}
}

func GetStudentRegistrationOutstandingData() (data []map[string]interface{}, err error) {
	mongodbStudentRegistrationRepository.Filter = bson.M{
		"$and": bson.A{
			bson.M{"status": "send"},
			bson.M{"admision": bson.M{"$exists": false}},
		},
	}
	data, err = mongodbStudentRegistrationRepository.GetMongoDB()
	return
}

func ApproveStudentRegistrationOutstandingData(userPayload interface{}, payload map[string]interface{}) (err error) {
	idPayload, _ := payload["_id"].(string)
	if pkg.IsEmptyString(idPayload) {
		err = errors.New("please fill id")
		return
	}
	id, _ := primitive.ObjectIDFromHex(idPayload)
	mongodbStudentRegistrationRepository.Filter = bson.D{{Key: "_id", Value: id}}
	mongodbStudentRegistrationRepository.Payload = bson.M{"admision": userPayload, "status": "Approved"}
	err = mongodbStudentRegistrationRepository.UpdateMongoDB()
	return
}

func GetAllStudentRegistrationAuthData(payload interface{}) (data []map[string]interface{}, err error) {
	var userPayload model.UserUpdatePayload
	payloadMap, ok := payload.(map[string]interface{})
	if !ok {
		err = fmt.Errorf("payload is not of type map[string]interface{}")
		return
	}

	payloadBytes, err := json.Marshal(payloadMap)
	if err != nil {
		return
	}

	if err = json.Unmarshal(payloadBytes, &userPayload); err != nil {
		return
	}

	log.Println(userPayload)
	if userPayload.Role == "admin" {
		mongodbStudentRegistrationRepository.Filter = bson.M{}
		data, err = mongodbStudentRegistrationRepository.GetMongoDB()
		return
	} else {
		id, _ := primitive.ObjectIDFromHex(userPayload.ID)
		mongodbStudentRegistrationRepository.Filter = bson.M{"admision._id": id}
		data, err = mongodbStudentRegistrationRepository.GetMongoDB()
		return
	}
}

func GetStudentRegistrationDetailData(searchPayload map[string]interface{}) (data []map[string]interface{}, err error) {
	idPayload, ok := searchPayload["_id"].(string)
	if ok {
		id, _ := primitive.ObjectIDFromHex(idPayload)
		delete(searchPayload, "_id")
		mongodbStudentRegistrationRepository.Filter = bson.D{{Key: "_id", Value: id}}
	} else {
		mongodbStudentRegistrationRepository.Filter = searchPayload
	}

	data, err = mongodbStudentRegistrationRepository.GetMongoDB()
	return
}
