package database

import (
	"context"
	"github.com/Filiphasan/golang-minify-url/configs"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type MongoContext struct {
	client    *mongo.Client
	appConfig *configs.AppConfig
}

func NewMongoContext(client *mongo.Client, appConfig *configs.AppConfig) *MongoContext {
	return &MongoContext{client: client, appConfig: appConfig}
}

func (m *MongoContext) Ping(ctx context.Context) error {
	return m.client.Ping(ctx, nil)
}

func (m *MongoContext) GetCollection(collectionName string) *mongo.Collection {
	return m.client.Database(m.appConfig.Mongodb.Database).Collection(collectionName)
}

func (m *MongoContext) Close(ctx context.Context) {
	_ = m.client.Disconnect(ctx)
}
