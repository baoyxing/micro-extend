package logutils

import "github.com/natefinch/lumberjack"

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
