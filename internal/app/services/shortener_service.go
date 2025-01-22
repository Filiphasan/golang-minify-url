package services

import (
	"context"
	"errors"
	"fmt"
	"github.com/Filiphasan/golang-minify-url/configs"
	"github.com/Filiphasan/golang-minify-url/internal/app/caches"
	"github.com/Filiphasan/golang-minify-url/internal/app/entities"
	"github.com/Filiphasan/golang-minify-url/internal/app/models/dtos/request"
	"github.com/Filiphasan/golang-minify-url/internal/app/models/dtos/response"
	"github.com/Filiphasan/golang-minify-url/internal/app/models/result"
	"github.com/Filiphasan/golang-minify-url/internal/database"
	"github.com/Filiphasan/golang-minify-url/pkg/constants"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"time"
)

type ShortenerService struct {
	appConfig    *configs.AppConfig
	cache        caches.Cache
	mongoContext *database.MongoContext
	tokenService *TokenService
}

func NewShortenerService(appConfig *configs.AppConfig, cache caches.Cache, mongoContext *database.MongoContext, tokenService *TokenService) *ShortenerService {
	return &ShortenerService{
		appConfig:    appConfig,
		cache:        cache,
		mongoContext: mongoContext,
		tokenService: tokenService,
	}
}

func (s *ShortenerService) GetShortenedUrl(ctx context.Context, r *request.GetShortenedURLRequest) result.HttpResult[*response.GetShortenedURLResponse] {
	token := r.Token
	cacheKey := fmt.Sprintf(constants.ShortUrlCacheKey, token)
	cacheValue, err := s.cache.Get(ctx, cacheKey)
	if err != nil && !errors.Is(err, redis.Nil) {
		return result.Error[*response.GetShortenedURLResponse](err)
	}

	res := &response.GetShortenedURLResponse{}
	if cacheValue != "" {
		res.LongUrl = cacheValue
		return result.Success[*response.GetShortenedURLResponse](res)
	}

	urlShorten := &entities.UrlShorten{}
	filter := bson.D{{"token", token}}
	err = s.mongoContext.UrlTokens().FindOne(ctx, filter).Decode(urlShorten)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return result.NotFound[*response.GetShortenedURLResponse]("Shortened URL not found!")
	}

	if err != nil {
		return result.Error[*response.GetShortenedURLResponse](err)
	}

	res.LongUrl = urlShorten.Url
	return result.Success[*response.GetShortenedURLResponse](res)
}

func (s *ShortenerService) SetShortenedUrl(ctx context.Context, r *request.SetShortenURLRequest) result.HttpResult[*response.SetShortenURLResponse] {
	token := ""
	url := r.Url

	urlShorten := &entities.UrlShorten{}
	urlExistErr := s.mongoContext.UrlTokens().FindOne(ctx, bson.D{{"url", url}}).Decode(urlShorten)
	if errors.Is(urlExistErr, mongo.ErrNoDocuments) {
		tokenRes := s.tokenService.GetUnusedToken(ctx)
		if tokenRes.Error != nil {
			return result.Error[*response.SetShortenURLResponse](tokenRes.Error)
		}
		token = tokenRes.Data
		expiredDay := r.ExpireDay
		urlShorten = entities.NewUrlShorten(token, url, expiredDay)
		_, err := s.mongoContext.UrlTokens().InsertOne(ctx, urlShorten)
		if err != nil {
			return result.Error[*response.SetShortenURLResponse](err)
		}

		cacheKey := fmt.Sprintf(constants.ShortUrlCacheKey, token)
		err = s.cache.Set(ctx, cacheKey, url, time.Hour*time.Duration(expiredDay))
		if err != nil {
			return result.Error[*response.SetShortenURLResponse](err)
		}
	} else if urlExistErr != nil {
		return result.Error[*response.SetShortenURLResponse](urlExistErr)
	} else {
		token = urlShorten.Token
	}

	sUrl := fmt.Sprintf("%s://%s:%s/%s", s.appConfig.Server.Scheme, s.appConfig.Server.Host, s.appConfig.Server.Port, token)
	res := &response.SetShortenURLResponse{
		Token:        token,
		ShortenedUrl: sUrl,
	}
	return result.Success[*response.SetShortenURLResponse](res)
}
