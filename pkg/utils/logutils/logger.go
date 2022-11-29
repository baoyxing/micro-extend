package logutils

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap/zapcore"
)

func NewLogger(opts ...Option) *lumberjack.Logger {
	config := DefaulutLumberjackConfig()
	for _, opt := range opts {
		opt.apply(config)
	}
	return &lumberjack.Logger{
		Filename:   config.filename,
		MaxSize:    config.maxSize,
		MaxAge:     config.maxAge,
		MaxBackups: config.maxBackups,
		LocalTime:  config.localTime,
		Compress:   config.compress,
	}
}

func NewZapEncoderConfig() *zapcore.EncoderConfig {
	return &zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "linenum",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,  // 小写编码器
		EncodeTime:     zapcore.ISO8601TimeEncoder,     // ISO8601 UTC 时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder, //
		EncodeCaller:   zapcore.FullCallerEncoder,      // 全路径编码器
		EncodeName:     zapcore.FullNameEncoder,
	}
}
