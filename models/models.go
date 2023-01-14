package models

import (
	"context"
	"errors"
	"fmt"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client      *mongo.Client
	ErrNotFound = errors.New("bootcamp not found")
)

func ConnectDB() *mongo.Client {
	var err error

	uri := viper.GetString("MONGO_URI")
	fmt.Println("Connecting to MongoDB...")

	// Set client options
	clientOptions := options.Client().ApplyURI(uri)

	client, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		panic(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Connected to MongoDB")
	}

	return client
}
