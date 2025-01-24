package logger

import (
	"github.com/Filiphasan/golang-minify-url/configs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var Logger *zap.Logger

func UseLogger(appConfig *configs.AppConfig) {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "@timestamp"
	encoderConfig.LevelKey = "LogLevel"
	encoderConfig.MessageKey = "Message"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	environment := appConfig.Environment

	core := zapcore.NewCore(zapcore.NewJSONEncoder(encoderConfig), zapcore.AddSync(os.Stdout), zap.InfoLevel)

	logger := zap.New(core)
	logger = logger.With(zap.String("ProjectName", appConfig.ProjectName), zap.String("Environment", environment))

	zap.ReplaceGlobals(logger)
	zap.RedirectStdLog(logger)

	Logger = logger
}
