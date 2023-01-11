package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	ZapLogEncoderConfig zapcore.EncoderConfig
	ZapLogEncoderJSON   zapcore.Encoder
	ZapLogEncoderHuman  zapcore.Encoder

	Logger zap.SugaredLogger // Global logger
)

// Use init() function in order to guarantee the logger is created every* time
func init() {

	ZapLogEncoderConfig = zap.NewProductionEncoderConfig()
	ZapLogEncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	ZapLogEncoderConfig.TimeKey = "@timestamp"
	ZapLogEncoderConfig.CallerKey = "debug.caller"
	ZapLogEncoderConfig.FunctionKey = "debug.function"
	ZapLogEncoderConfig.MessageKey = "message"

	ZapLogEncoderJSON = zapcore.NewJSONEncoder(ZapLogEncoderConfig)
	ZapLogEncoderHuman = zapcore.NewConsoleEncoder(ZapLogEncoderConfig)

	Logger = *zap.New(
		zapcore.NewCore(ZapLogEncoderHuman, zapcore.AddSync(os.Stdout), zap.WarnLevel),
	).Sugar()
}

func ResetLevel(logLevel zapcore.Level) zap.SugaredLogger {
	Logger = *zap.New(
		zapcore.NewCore(ZapLogEncoderHuman, zapcore.AddSync(os.Stdout), logLevel),
	).Sugar()
	return Logger
}

func ResetEncoder(encoder zapcore.Encoder) zap.SugaredLogger {
	Logger = *zap.New(
		zapcore.NewCore(ZapLogEncoderHuman, zapcore.AddSync(os.Stdout), Logger.Level()),
	).Sugar()
	return Logger
}
