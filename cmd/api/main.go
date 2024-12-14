package main

import (
	"github.com/Filiphasan/golang-minify-url/configs"
	"github.com/Filiphasan/golang-minify-url/internal/logger"
	"go.uber.org/zap"
)

func main() {
	configs.LoadAppConfig()
	appConfig := configs.GetAppConfig()

	logger.UseLogger(appConfig)
	loggerErr := logger.Logger.Sync()
	if loggerErr != nil {
		panic(loggerErr)
	}

	logger.Logger.Info("Hello, World!", zap.String("Owner", "Hasan Erdal"))
}
