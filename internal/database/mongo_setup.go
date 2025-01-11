package database

import (
	"fmt"
	"github.com/Filiphasan/golang-minify-url/configs"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func UseMongo(appConfig *configs.AppConfig) *mongo.Client {
	uri := getMongoUri(appConfig)
	clientOptions := options.Client().ApplyURI(uri)
	clientOptions.Auth = &options.Credential{Username: appConfig.Mongodb.Username, Password: appConfig.Mongodb.Password}
	client, err := mongo.Connect(clientOptions)
	if err != nil {
		panic(err)
	}

	return client
}

func getMongoUri(appConfig *configs.AppConfig) string {
	return fmt.Sprintf("mongodb://%s:%s", appConfig.Mongodb.Host, appConfig.Mongodb.Port)
}
