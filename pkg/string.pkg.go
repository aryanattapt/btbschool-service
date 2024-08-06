package pkg

import (
	"context"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func IsEmptyString(data string) bool {
	return len(strings.TrimSpace(data)) == 0 || strings.TrimSpace(data) == ""
}

func GenerateRandomString(length int) string {
	var randomizer = rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
	var letters = []rune("abcdefghijklmnopqrstuvwxyz1234567890ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, length)
	for i := range b {
		b[i] = letters[randomizer.Intn(len(letters))]
	}
	return string(b)
}

func GenerateRandomNumber(length int) string {
	var randomizer = rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
	var letters = []rune("1234567890")
	b := make([]rune, length)
	for i := range b {
		b[i] = letters[randomizer.Intn(len(letters))]
	}
	return string(b)
}

func GenerateUUID() string {
	return uuid.New().String()
}

type counter struct {
	Module  string `bson:"module,omitempty" json:"module,omitempty"`
	Date    string `bson:"date,omitempty" json:"date,omitempty"`
	Counter int    `bson:"counter,omitempty" json:"counter,omitempty"`
}

func GetFormattedCounter(prefix string, module string) (result string, err error) {
	ctx := context.TODO()
	mongoDatabase := MongoDBDatabase{DatabaseName: "btb_app", CollectionName: "counter"}
	db, err := mongoDatabase.ConnectMongoDB()
	if err != nil {
		return
	}
	defer mongoDatabase.DisconnectMongoDB(db.Client())

	collection := db.Collection(mongoDatabase.CollectionName)
	today := time.Now().Format("02012006")

	filter := bson.M{"date": today, "module": module}
	update := bson.M{
		"$setOnInsert": bson.M{"date": today},
		"$inc":         bson.M{"counter": 1},
	}

	var resultObj counter
	opts := options.FindOneAndUpdate().SetUpsert(true).SetReturnDocument(options.After)
	err = collection.FindOneAndUpdate(ctx, filter, update, opts).Decode(&resultObj)
	if err != nil {
		return
	}
	result = fmt.Sprintf("%s%s%03d", prefix, today, resultObj.Counter)
	return
}
