package database

import (
	"context"
	"github.com/Filiphasan/golang-minify-url/configs"
	"github.com/Filiphasan/golang-minify-url/internal/app/entities"
	"github.com/Filiphasan/golang-minify-url/pkg/constants"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"sync"
)

var once sync.Once

type MongoContext struct {
	client    *mongo.Client
	appConfig *configs.AppConfig
	Database  *mongo.Database
}

func NewMongoContext(client *mongo.Client, appConfig *configs.AppConfig) *MongoContext {
	database := client.Database(appConfig.Mongodb.Database)
	return &MongoContext{client: client, appConfig: appConfig, Database: database}
}

func (m *MongoContext) EnsureIndexes(ctx context.Context) {
	once.Do(func() {
		urlTokenCollection := m.UrlTokens()
		_, _ = urlTokenCollection.Indexes().CreateMany(ctx, entities.EnsureUrlTokenIndex())

		urlShortenCollection := m.UrlShortens()
		_, _ = urlShortenCollection.Indexes().CreateMany(ctx, entities.EnsureUrlShortenIndex())
	})
}

func (m *MongoContext) Ping(ctx context.Context) error {
	err := m.client.Ping(ctx, nil)
	return err
}

func (m *MongoContext) Close(ctx context.Context) {
	_ = m.client.Disconnect(ctx)
}

// Collection methods

func (m *MongoContext) UrlTokens() *mongo.Collection {
	collection := m.Database.Collection(constants.UrlToken)
	return collection
}

func (m *MongoContext) UrlShortens() *mongo.Collection {
	collection := m.Database.Collection(constants.UrlShorten)
	return collection
}
