package chat

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MessageRepo interface {
	FindAll() []Msg
	Save(msg *Msg) error
}

type MongoMessageRepo struct {
	client *mongo.Client
}

func (repo *MongoMessageRepo) FindAll() []Msg {
	var res []Msg

	coll := repo.client.Database("chat").Collection("messages")
	cursor, err := coll.Find(context.TODO(), bson.D{})
	if err != nil {
		return nil
	}

	err = cursor.All(context.TODO(), &res)
	if err != nil {
		return nil
	}

	return res
}

func (repo *MongoMessageRepo) Save(msg *Msg) error {
	coll := repo.client.Database("chat").Collection("messages")
	_, err := coll.InsertOne(context.TODO(), msg)
	if err != nil {
		return err
	}

	return nil
}
