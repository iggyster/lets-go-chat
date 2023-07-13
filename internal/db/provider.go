package db

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ProvideClient() (*mongo.Client, func(), error) {
	uri := os.Getenv("MONGO_URI")
	opts := options.Client().ApplyURI(uri)

	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("mongo db connection is established")

	return client, func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}, nil
}
