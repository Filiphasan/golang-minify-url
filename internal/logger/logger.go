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
	encoderConfig.TimeKey = "timestamp"
	encoderConfig.LevelKey = "logLevel"
	encoderConfig.MessageKey = "message"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	core := zapcore.NewCore(zapcore.NewJSONEncoder(encoderConfig), zapcore.AddSync(os.Stdout), zap.InfoLevel)

	logger := zap.New(core)
	logger = logger.With(zap.String("projectName", appConfig.ProjectName))

	zap.ReplaceGlobals(logger)

	Logger = logger
}
