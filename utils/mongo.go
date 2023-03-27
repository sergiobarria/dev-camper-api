package utils

import (
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	db := viper.GetString("DATABASE")

	return client.Database(db).Collection(collectionName)
}
