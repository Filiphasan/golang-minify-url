package main

import (
	"github.com/Filiphasan/golang-minify-url/configs"
	"github.com/Filiphasan/golang-minify-url/internal/logger"
	"github.com/Filiphasan/golang-minify-url/internal/setup"
	"github.com/gin-gonic/gin"
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

	router := gin.Default()
	setup.NewApp(router).SetupApp()

	logger.Logger.Info("Hello, Golang Minify URL!", zap.String("Owner", "Hasan Erdal"))
	_ = router.Run(":5001")
}
