package logger

import (
	"github.com/zeromicro/go-zero/core/logx"
	"testing"
)

func TestZap(t *testing.T) {
	zapLoggerConfig := Config{
		Level:         "debug",
		Stacktrace:    true,
		AddCaller:     true,
		CallerShip:    1,
		Mode:          "console",
		FileName:      "tetest",
		ErrorFileName: "tetest",
		MaxSize:       0,
		MaxAge:        0,
		MaxBackup:     0,
		Async:         false,
		Json:          false,
		Compress:      false,
		options:       nil,
	}
	InitLogger(zapLoggerConfig)

	logx.SetWriter(NewZapWriter(L))
	logx.Infow("111", logx.Field("test", "test"))
	logx.Debugw("111", logx.Field("test", "test"))
	logx.Slowf("111")
	logx.Errorw("111")
}
