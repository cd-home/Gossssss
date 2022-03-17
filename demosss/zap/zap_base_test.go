package zap

import (
	"testing"

	"go.uber.org/zap"
)

func TestSugarZap(t *testing.T) {
	//  NewProduction => NewProductionConfig => info level json
	// logger, _ := zap.NewProduction()

	// NewDevelopment => NewDevelopmentConfig => debug level, console
	logger, _ := zap.NewDevelopment()

	// Sync calls the underlying Core's Sync method, flushing any buffered log
	// entries. Applications should take care to call Sync before exiting.
	defer logger.Sync()

	sugar := logger.Sugar()

	// args mode
	sugar.Info("info")

	// k-v
	sugar.Infow("info", "url", "/info", "attemps", 3)

	// format
	sugar.Infof("url %s", "/info")
}

func TestNewProductionZap(t *testing.T) {
	// we should use the production
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	// args type more strict than sugar
	// suport zap.Field
	logger.Info("info", zap.String("module", "zap"), zap.Any("any", []int{1, 2, 3}))
}
