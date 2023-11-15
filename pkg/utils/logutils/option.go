package logutils

import (
	"go.uber.org/zap/zapcore"
	"time"
)

const (
	DefaultLogPath = "./log"
)

type Option interface {
	apply(cfg *OptionConfig)
}

type option func(cfg *OptionConfig)

func (fn option) apply(cfg *OptionConfig) {
	fn(cfg)
}

type OptionConfig struct {
	path             string        // 日志文件路径，默认 ./log
	maxSize          int           // 每个日志文件保存10M，默认 100M
	maxAge           int           // 保留30个备份，默认不限
	maxBackups       int           // 保留7天，默认不限
	compress         bool          //是否压缩
	outputMode       int           ////输出模式 1:控制台 2：文件 3：控制台和文件都输出
	rotationDuration time.Duration // 最大分隔时长 单位小时
	suffix           string        //后缀名称 默认.log
}

func DefaultOptionConfig() *OptionConfig {
	return &OptionConfig{
		path:             DefaultLogPath,
		maxSize:          10,
		maxAge:           7,
		maxBackups:       30,
		compress:         true,
		outputMode:       3,
		rotationDuration: time.Duration(24),
		suffix:           ".log",
	}
}

func WithPath(path string) Option {
	return option(func(cfg *OptionConfig) {
		if path != "" {
			cfg.path = path
		}

	})
}

func WithMaxSize(maxSize int) Option {
	return option(func(cfg *OptionConfig) {
		if maxSize > 0 {
			cfg.maxSize = maxSize
		}

	})
}

func WithMaxAge(maxAge int) Option {
	return option(func(cfg *OptionConfig) {
		if maxAge > 0 {
			cfg.maxAge = maxAge
		}
	})
}

func WithMaxBackups(maxBackups int) Option {
	return option(func(cfg *OptionConfig) {
		if maxBackups > 0 {
			cfg.maxBackups = maxBackups
		}
	})
}

func WithCompress(compress bool) Option {
	return option(func(cfg *OptionConfig) {
		cfg.compress = compress
	})
}

func WithOutputMode(outputMode int) Option {
	return option(func(cfg *OptionConfig) {
		cfg.outputMode = outputMode
	})
}

func WithRotationDuration(rotationDuration time.Duration) Option {
	return option(func(cfg *OptionConfig) {
		cfg.rotationDuration = rotationDuration
	})
}
func WithSuffix(suffix string) Option {
	return option(func(cfg *OptionConfig) {
		cfg.suffix = suffix
	})
}

// encoderConfig copy from hlogzap
func pkgEncoderConfig() zapcore.EncoderConfig {
	cfg := coreEncoderConfig()
	cfg.EncodeTime = ISO8601TimeEncoder
	cfg.EncodeLevel = CapitalLevelEncoder
	cfg.EncodeDuration = StringDurationEncoder
	cfg.EncodeCaller = FullCallerEncoder
	return cfg
}

func FullCallerEncoder(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
	// TODO: consider using a byte-oriented API to save an allocation.
	enc.AppendString("[" + caller.String() + "]")
}

// ISO8601TimeEncoder 自定义时间格式
func ISO8601TimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	encodeTimeLayout(t, "2006-01-02 15:04:05", enc)
}

func encodeTimeLayout(t time.Time, layout string, enc zapcore.PrimitiveArrayEncoder) {
	type appendTimeEncoder interface {
		AppendTimeLayout(time.Time, string)
	}

	if enc, ok := enc.(appendTimeEncoder); ok {
		enc.AppendTimeLayout(t, layout)
		return
	}
	enc.AppendString("[" + t.Format(layout) + "]")
}

// CapitalLevelEncoder 自定义等级格式
func CapitalLevelEncoder(l zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString("[" + l.CapitalString() + "]")
}

// StringDurationEncoder 自定义时间格式
func StringDurationEncoder(d time.Duration, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString("[" + d.String() + "]")
}

func ShortCallerEncoder(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString("[" + caller.TrimmedPath() + "]")
}
