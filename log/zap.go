package log

import (
	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const logSkip = 1

func retZap() Interface {
	config := zap.NewProductionConfig()
	config.EncoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder
	config.Encoding = "json"
	ZapLog, err := config.Build(zap.AddCallerSkip(logSkip))
	if err != nil {
		panic(fmt.Sprintf("zap log: %v", err))
	}

	return ZapLog.Sugar()
}
