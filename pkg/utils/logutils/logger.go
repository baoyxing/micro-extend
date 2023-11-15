package logutils

import (
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap/zapcore"
	"log"
	"time"
)

func NewWriteSyncer(fileName string, config *OptionConfig) zapcore.WriteSyncer {
	if config.rotationDuration > 0 {
		return NewRotateWriteSyncer(fileName, config)
	} else {
		return NewLumberjackWriteSyncer(fileName, config)
	}
}

func NewLumberjackWriteSyncer(fileName string,
	config *OptionConfig) zapcore.WriteSyncer {
	return zapcore.AddSync(&lumberjack.Logger{
		Filename:   config.path + fileName + config.suffix,
		MaxSize:    config.maxSize,
		MaxAge:     config.maxAge,
		MaxBackups: config.maxBackups,
		Compress:   config.compress,
	})

}

func NewRotateWriteSyncer(fileName string,
	config *OptionConfig) zapcore.WriteSyncer {
	fileName = config.path + fileName
	logf, err := rotatelogs.New(
		fileName+"_%Y%m%d%H%M"+config.suffix,
		rotatelogs.WithLinkName(fileName+config.suffix),
		rotatelogs.WithMaxAge(time.Duration(config.maxAge)*24*time.Hour),
		rotatelogs.WithRotationTime(config.rotationDuration*time.Minute),
	)
	if err != nil {
		log.Printf("failed to create rotatelogs: %s", err)
	}
	return zapcore.AddSync(logf)
}

func coreEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		TimeKey:        "ts",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,  // 小写编码器
		EncodeTime:     zapcore.EpochTimeEncoder,       // ISO8601 UTC 时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder, //
		EncodeCaller:   ShortCallerEncoder,             // 全路径编码器
		EncodeName:     zapcore.FullNameEncoder,
	}
}
