package database

import (
	"github.com/Filiphasan/golang-minify-url/configs"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type MongoContext struct {
	client *mongo.Client
}

func NewMongoContext(client *mongo.Client, appConfig *configs.AppConfig) *MongoContext {
	return &MongoContext{client: client}
}

func (m *MongoContext) GetCollection(collectionName string) *mongo.Collection {
	return m.client.Database("golang-minify-url").Collection(collectionName)
}
