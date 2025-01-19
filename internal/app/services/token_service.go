package services

import (
	"context"
	"errors"
	"github.com/Filiphasan/golang-minify-url/configs"
	"github.com/Filiphasan/golang-minify-url/internal/app/caches"
	"github.com/Filiphasan/golang-minify-url/internal/app/entities"
	"github.com/Filiphasan/golang-minify-url/internal/app/models/result"
	"github.com/Filiphasan/golang-minify-url/internal/database"
	"github.com/Filiphasan/golang-minify-url/pkg/constants"
	"github.com/Filiphasan/golang-minify-url/pkg/helpers"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type TokenService struct {
	appConfig    *configs.AppConfig
	cache        caches.Cache
	mongoContext *database.MongoContext
}

func NewTokenService(appConfig *configs.AppConfig, cache caches.Cache, mongoContext *database.MongoContext) *TokenService {
	return &TokenService{
		appConfig:    appConfig,
		cache:        cache,
		mongoContext: mongoContext,
	}
}

func (t *TokenService) GenerateAndSaveToken(ctx context.Context, tokenBuilder *helpers.TokenBuilder) result.HttpResult[string] {
	if tokenBuilder == nil {
		tokenBuilder = helpers.NewTokenBuilder().
			SetEpoch(t.appConfig.Token.EpochDate).
			SetAddChars(3)
	}

	token := tokenBuilder.Build()
	_, err := t.mongoContext.UrlTokens().InsertOne(ctx, entities.NewUrlToken(token))
	if err != nil {
		return result.Error[string](err)
	}
	return result.Success[string](token)
}

func (t *TokenService) GetUnusedToken(ctx context.Context) result.HttpResult[string] {
	token, err := t.cache.ListPop(ctx, constants.TokenSeedListCacheKey, false)
	if err != nil && !errors.Is(err, redis.Nil) {
		return result.Error[string](err)
	}

	if token != "" {
		err = t.setTokenUsed(ctx, token)
		if err != nil {
			return result.Error[string](err)
		}
		return result.Success[string](token)
	}

	res := t.GenerateAndSaveToken(ctx, nil)
	if res.Error != nil {
		return res
	}

	token = res.Data
	err = t.setTokenUsed(ctx, token)
	if err != nil {
		return result.Error[string](err)
	}
	return result.Success[string](token)
}

func (t *TokenService) setTokenUsed(ctx context.Context, token string) error {
	filter := bson.D{{"token", token}}
	_, err := t.mongoContext.UrlTokens().UpdateOne(ctx, filter, bson.M{"$set": bson.M{"used": true}})
	if err != nil {
		return err
	}
	return nil
}
