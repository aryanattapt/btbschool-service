package repository

import (
	"btb-service/model"
	"btb-service/pkg"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	mongoDBUserRepository = pkg.MongoDBDatabase{DatabaseName: "btb_client", CollectionName: "users"}
)

func GetUserByUsernameOrEmail(username string, email string) (data []map[string]interface{}, err error) {
	mongoDBUserRepository.Filter = bson.M{
		"isactive": "active",
		"$or": bson.A{
			bson.M{"username": username},
			bson.M{"email": email},
		},
	}
	data, err = mongoDBUserRepository.GetMongoDB()
	return
}

func GetUserById(userid string) (data []map[string]interface{}, err error) {
	id, _ := primitive.ObjectIDFromHex(userid)
	mongoDBUserRepository.Filter = bson.M{
		"isactive": "active",
		"_id":      id,
	}
	data, err = mongoDBUserRepository.GetMongoDB()
	return
}

func SaveUser(payload model.UserInsertPayload) (err error) {
	payload.RegisteredDate = primitive.NewDateTimeFromTime(time.Now())
	payload.IsActive = "active"
	data, _ := pkg.StructToMap(payload)
	mongoDBUserRepository.Payload = data
	err = mongoDBUserRepository.InsertMongoDB()
	return
}

func UpdateUser(payload model.UserUpdatePayload) (err error) {
	id, _ := primitive.ObjectIDFromHex(payload.ID)
	mongoDBUserRepository.Filter = bson.M{"_id": id}

	payload.UpdatedDate = primitive.NewDateTimeFromTime(time.Now())
	mongoDBUserRepository.Payload = bson.M{
		"firstname":   payload.FirstName,
		"lastname":    payload.LastName,
		"username":    payload.Username,
		"email":       payload.Email,
		"role":        payload.Role,
		"updateddate": payload.UpdatedDate,
		"isactive":    payload.IsActive,
		"permission":  payload.Permission,
	}

	if !pkg.IsEmptyString(payload.Password) {
		mongoDBUserRepository.Payload["password"] = pkg.HashPasswordBCrypt(payload.Password)
	}

	err = mongoDBUserRepository.UpdateMongoDB()
	return
}

func GetAllUser(searchPayload map[string]interface{}) (data []map[string]interface{}, err error) {
	idPayload, ok := searchPayload["_id"].(string)
	if ok {
		id, _ := primitive.ObjectIDFromHex(idPayload)
		delete(searchPayload, "_id")
		mongoDBUserRepository.Filter = bson.D{{Key: "_id", Value: id}}
	} else {
		mongoDBUserRepository.Filter = searchPayload
	}
	data, err = mongoDBUserRepository.GetMongoDB()
	return
}

func CheckPermission(userid string, permission string) (err error) {
	id, _ := primitive.ObjectIDFromHex(userid)
	mongoDBUserRepository.Filter = bson.M{
		"isactive":   "active",
		"_id":        id,
		"permission": permission,
	}

	data, err := mongoDBUserRepository.GetMongoDB()
	if len(data) == 0 {
		err = errors.New("unauthorized")
	}
	return
}
