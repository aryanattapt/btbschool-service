package repository

import (
	"btb-service/pkg"
	"errors"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var mongoDBIgRepository = pkg.MongoDBDatabase{DatabaseName: "btb_app", CollectionName: "token"}

func GetInstagramToken() (data string, err error) {
	mongoDBIgRepository.Filter = bson.M{"type": "instagramtoken"}
	queryData, err := mongoDBIgRepository.GetMongoDB()

	if len(queryData) > 0 {
		data, _ = queryData[0]["token"].(string)
	} else {
		err = errors.New("token Not Found")
	}
	return
}

func UpdateInstagramToken(token string) (err error) {
	log.Println("Update new instagram token : ", token)
	mongoDBIgRepository.Payload = bson.M{"token": token, "updateddate": primitive.NewDateTimeFromTime(time.Now())}
	mongoDBIgRepository.Filter = bson.M{"type": "instagramtoken"}
	err = mongoDBIgRepository.UpdateMongoDB()
	return
}
