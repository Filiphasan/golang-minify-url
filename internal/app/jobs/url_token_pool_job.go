package jobs

import (
	"context"
	"github.com/Filiphasan/golang-minify-url/configs"
	"github.com/Filiphasan/golang-minify-url/internal/app/caches"
	"github.com/Filiphasan/golang-minify-url/internal/app/entities"
	"github.com/Filiphasan/golang-minify-url/internal/database"
	"github.com/Filiphasan/golang-minify-url/pkg/constants"
	"github.com/Filiphasan/golang-minify-url/pkg/helpers"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.uber.org/zap"
	"sync"
)

type UrlTokenPoolJob struct {
	appConfig    *configs.AppConfig
	logger       *zap.Logger
	mongoContext *database.MongoContext
	cache        caches.Cache
}

func NewUrlTokenPoolJob(appConfig *configs.AppConfig, logger *zap.Logger, mongoContext *database.MongoContext, cache caches.Cache) *UrlTokenPoolJob {
	return &UrlTokenPoolJob{appConfig: appConfig, logger: logger, mongoContext: mongoContext, cache: cache}
}

func (u *UrlTokenPoolJob) Run() {
	const methodName = "UrlTokenPoolJob.Run"
	const MaxDegreeOfParallelism = 10

	u.logger.Info("Token pool job is running", zap.String("Method", methodName))
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	unusedFilter := bson.D{{"isUsed", false}}
	unusedCount, err := u.mongoContext.UrlTokens().CountDocuments(ctx, unusedFilter)
	if err != nil {
		u.logger.Error("Failed to count documents", zap.Error(err), zap.String("Method", methodName))
		return
	}

	if unusedCount >= int64(u.appConfig.Token.PoolingSize) {
		u.logger.Info("Token pool is already full", zap.String("Method", methodName))
		return
	}

	tokenBuilder := helpers.NewTokenBuilder().
		SetEpoch(u.appConfig.Token.EpochDate).
		SetAddChars(4)
	// Calculate the number of tokens to be generated
	remainingTokens := u.appConfig.Token.PoolingSize - int(unusedCount) + u.appConfig.Token.ExtendSize

	tasks := make(chan int, remainingTokens)
	var wg sync.WaitGroup
	for i := 0; i < MaxDegreeOfParallelism; i++ {
		wg.Add(1)
		go u.GenerateAndSaveToken(ctx, &wg, tasks, tokenBuilder, methodName)
	}

	for i := 0; i < remainingTokens; i++ {
		tasks <- i
	}
	close(tasks)

	wg.Wait()
	u.logger.Info("Token pool job is done", zap.String("Method", methodName))
}

func (u *UrlTokenPoolJob) GenerateAndSaveToken(ctx context.Context, wg *sync.WaitGroup, tasks <-chan int, tokenBuilder *helpers.TokenBuilder, methodName string) {
	defer wg.Done()

	for range tasks {
		token := tokenBuilder.Build()
		urlToken := entities.NewUrlToken(token)
		_, err := u.mongoContext.UrlTokens().InsertOne(ctx, urlToken)
		if err != nil {
			u.logger.Error("Failed to insert token", zap.Error(err), zap.String("Method", methodName))
			return
		}
		err = u.cache.AddList(ctx, constants.TokenSeedListCacheKey, true, token)
		if err != nil {
			u.logger.Error("Failed to add token to cache", zap.Error(err), zap.String("Method", methodName))
			return
		}
	}
}
