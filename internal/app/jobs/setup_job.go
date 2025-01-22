package jobs

import (
	"github.com/Filiphasan/golang-minify-url/configs"
	"github.com/Filiphasan/golang-minify-url/internal/app/caches"
	"github.com/Filiphasan/golang-minify-url/internal/database"
	"github.com/go-co-op/gocron/v2"
	"go.uber.org/zap"
)

type SetupJob struct {
	appConfig    *configs.AppConfig
	logger       *zap.Logger
	mongoContext *database.MongoContext
	cache        caches.Cache
}

func NewSetupJob(appConfig *configs.AppConfig, logger *zap.Logger, mongoContext *database.MongoContext, cache caches.Cache) *SetupJob {
	return &SetupJob{appConfig: appConfig, logger: logger, mongoContext: mongoContext, cache: cache}
}

func (s *SetupJob) Run() {
	const methodName = "SetupJob.Run"

	scheduler, err := gocron.NewScheduler()
	if err != nil {
		s.logger.Error("Failed to create scheduler", zap.Error(err), zap.String("Method", methodName))
		return
	}

	// Add your jobs here
	urlTokenPool := NewUrlTokenPoolJob(s.appConfig, s.logger, s.mongoContext, s.cache)
	urlTokenPoolJob, err := scheduler.NewJob(
		gocron.CronJob("* */10 * * * *", true),
		gocron.NewTask(urlTokenPool.Run),
		gocron.WithName("UrlTokenPoolJob"),
	)
	if err != nil {
		s.logger.Error("Failed to create job", zap.Error(err), zap.String("JobName", urlTokenPoolJob.Name()), zap.String("Method", methodName))
		return
	}
}
