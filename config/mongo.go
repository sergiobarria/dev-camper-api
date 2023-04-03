package config

import (
	"context"

	"github.com/fatih/color"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Client // this is not used in the project, but it's a good example of how to use the mongo client as a global variable

func NewMongoClient() *mongo.Client {
	uri := EnvVars.MONGO_URI
	db := EnvVars.MONGO_DB

	serverAPI := options.ServerAPI(options.ServerAPIVersion1) // Set MongoDB server API version to 1
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)
	ctx, cancel := context.WithTimeout(context.Background(), 10*1000)
	defer cancel()

	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		panic(err)
	}

	var result bson.M
	if err := client.Database(db).RunCommand(context.Background(), bson.D{{Key: "ping", Value: 1}}).Decode(&result); err != nil {
		panic(err)
	}

	c := color.New(color.FgHiYellow).Add(color.Bold)
	c.Println("⇨ 💾 Successfully connected to MongoDB")

	return client
}

func GetCollection(client *mongo.Client, coll string) *mongo.Collection {
	return client.Database(viper.GetString("DATABASE")).Collection(coll)
}
