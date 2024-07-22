package repository

import (
	"btb-service/model"
	"btb-service/pkg"

	"go.mongodb.org/mongo-driver/bson"
)

var mongodbAlumniReposiroty = pkg.MongoDBDatabase{DatabaseName: "btb_client", CollectionName: "alumni"}

func GetAlumni() (data []map[string]interface{}, err error) {
	mongodbAlumniReposiroty.Filter = bson.D{{}}
	data, err = mongodbAlumniReposiroty.GetMongoDB()
	return
}

func SaveAlumni(payload model.AlumniInsertPayload) (err error) {
	data, _ := pkg.StructToMap(payload)
	mongodbAlumniReposiroty.Payload = data
	err = mongodbAlumniReposiroty.InsertMongoDB()
	return
}
