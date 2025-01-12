package logger_test

import (
	"github.com/accoladexin/vblog/common/logger"
	"testing"
	"time"
)

func TestLoggerDebug(t *testing.T) {
	// 同步编程
	//1.xxx
	time.Sleep(1 * time.Second)
	//2.xxx

	// js 是异步编程

	logger.L().Debug().Str("host", "10.10.10.1").Msg("hello zero log")
}
