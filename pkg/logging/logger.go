package logging

import (
	"fmt"
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-ybaspinar/pkg/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewLogger(config *config.Config) {
	loglevel, err := zapcore.ParseLevel(config.Logger.Level)
	if err != nil {
		panic(fmt.Sprintf("Unknown log level %v", loglevel))
	}
	var cfg zap.Config
	if config.Logger.Development {
		cfg = zap.NewDevelopmentConfig()
	} else {
		cfg = zap.NewProductionConfig()
	}
	logger, err := cfg.Build()
	if err != nil {
		logger = zap.NewNop()
	}
	zap.ReplaceGlobals(logger)
}
func Close() {
	defer zap.L().Sync()
}
