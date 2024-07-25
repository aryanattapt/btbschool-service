package repository

import (
	"btb-service/model"
	"btb-service/pkg"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var mongodbStudentRegistrationRepository = pkg.MongoDBDatabase{DatabaseName: "btb_client", CollectionName: "studentregistration"}

func GetStudentRegistrationOutstandingData(payload model.StudentRegistrationOutstandingDataPayload) (data []map[string]interface{}, err error) {
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
