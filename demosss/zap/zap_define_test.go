package zap

import (
	"io"
	"os"
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
	// e. Extending zap to support a new encoding (e.g., BSON), a new log sink (e.g., Kafka)

	highPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.WarnLevel
	})

	lowPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.InfoLevel
	})

	kafkaEncode := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
	// Discard is an io.Writer on which all Write calls succeed without doing anything.
	topicDebugging := zapcore.AddSync(io.Discard)

	consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
	consoleDebugging := zapcore.Lock(os.Stdout)

	var cores []zapcore.Core

	// all core
	cores = append(cores, zapcore.NewCore(kafkaEncode, topicDebugging, highPriority))
	cores = append(cores, zapcore.NewCore(consoleEncoder, consoleDebugging, lowPriority))
	core := zapcore.NewTee(cores...)

	log := zap.New(core)
	defer log.Sync()
	log = log.WithOptions(zap.Fields(zap.Field{Key: "module", String: "user", Type: zapcore.StringType}))

	// set highPriority or lowPriority, while execute info debug warn error will check
	log.Debug("debug", zap.String("test", "tes2"))
	log.Info("info", zap.String("test", "tes2"))
	log.Warn("warn", zap.String("test", "tes2"))
	log.Error("error", zap.String("test", "tes2"))
}
