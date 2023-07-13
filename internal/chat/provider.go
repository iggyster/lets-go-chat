package chat

import (
	"github.com/google/wire"
	"go.mongodb.org/mongo-driver/mongo"
)

var MongoMessageRepoProviderSet = wire.NewSet(
	ProvideMessageRepo,
	wire.Bind(new(MessageRepo), new(*MongoMessageRepo)),
)

func ProvideMessageRepo(client *mongo.Client) *MongoMessageRepo {
	return &MongoMessageRepo{client: client}
}
