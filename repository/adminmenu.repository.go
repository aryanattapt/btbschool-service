package repository

import (
	"btb-service/pkg"
	"log"

	"go.mongodb.org/mongo-driver/bson"
)

var (
	mongoDBAdminMenuRepository = pkg.MongoDBDatabase{DatabaseName: "btb_app", CollectionName: "adminmenus"}
)

func GetAdminMenu(permission []string) (data []map[string]interface{}, err error) {
	log.Println(permission)
	mongoDBAdminMenuRepository.Filter = bson.M{
		"$or": []bson.M{
			{
				"items.auth": bson.M{
					"$in": permission,
				},
			},
			{
				"auth": bson.M{
					"$in": permission,
				},
			},
		},
	}

	data, err = mongoDBAdminMenuRepository.GetMongoDB()
	return
}
