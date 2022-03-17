package zap

import (
	"testing"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func TestZapFlow(t *testing.T) {
	// step1: logger config and encoderConfig
	config := zap.Config{
		Level:         zap.NewAtomicLevelAt(zap.InfoLevel),
		InitialFields: map[string]interface{}{"serviceName": "spikeProxy"},
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:       "time",
			LevelKey:      "level",
			NameKey:       "logger",
			CallerKey:     "caller",
			FunctionKey:   zapcore.OmitKey,
			MessageKey:    "msg",
			StacktraceKey: "stacktrace",
			LineEnding:    zapcore.DefaultLineEnding,
			EncodeLevel:   zapcore.LowercaseLevelEncoder,
			EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
				enc.AppendString(t.Local().Format("2006-01-02 15:04:05"))
			},
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
		Encoding:         "json",
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
	}
	// step2: build config
	// build encoder build sink(WriteSyncer)
	// zap.NewCore
	// we can DIY config and sink(WriteSyncer) to output files or kafka
	logger, _ := config.Build()

	// step3: add some extro encoder
	logger = logger.WithOptions(
		zap.AddCaller(),
		// common extro fileds
		zap.Fields(zapcore.Field{Key: "module", String: "user", Type: zapcore.StringType}),
	)

	// special
	logger.Info("info", zap.String("some", "above"))
	logger.Debug("debug")
	logger.Warn("warn")
	logger.Error("error")
}

func TestZapDefinePractice(t *testing.T) {
	// highPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
	// 	return lvl >= zapcore.ErrorLevel
	// })

	// fmt.Printf("%T\n", highPriority)
	// lowPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
	// 	return lvl < zapcore.ErrorLevel
	// })

	// kafkaEncode := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())

	// zapcore.NewCore(kafkaEncode, lowPriority)
	// zapcore.NewTee()
	// log, _ := zap.NewProduction()
	// defer log.Sync()
	// log = log.WithOptions(zap.Fields(zap.Field{Key: "module", String: "user", Type: zapcore.StringType}))
	// log.Info("999", zap.String("test", "tes2"))
}
