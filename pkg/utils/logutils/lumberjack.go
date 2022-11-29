package logutils

const (
	DefalutLogPath = "app.log"
)

type Option interface {
	apply(cfg *LumberjackConfig)
}

type option func(cfg *LumberjackConfig)

func (fn option) apply(cfg *LumberjackConfig) {
	fn(cfg)
}

type LumberjackConfig struct {
	filename   string // 日志文件路径，默认 app.loh
	maxSize    int    // 每个日志文件保存10M，默认 100M
	maxAge     int    // 保留30个备份，默认不限
	maxBackups int    // 保留7天，默认不限
	compress   bool   //是否压缩
	localTime  bool   //本地
}

func DefaulutLumberjackConfig() *LumberjackConfig {
	return &LumberjackConfig{
		filename:   DefalutLogPath,
		maxSize:    10,
		maxAge:     7,
		maxBackups: 30,
		compress:   true,
		localTime:  true,
	}
}

func WithLumberjackFilename(filename string) Option {
	return option(func(cfg *LumberjackConfig) {
		cfg.filename = filename
	})
}

func WithLumberjackMaxSize(maxSize int) Option {
	return option(func(cfg *LumberjackConfig) {
		cfg.maxSize = maxSize
	})
}

func WithLumberjackMaxAge(maxAge int) Option {
	return option(func(cfg *LumberjackConfig) {
		cfg.maxAge = maxAge
	})
}

func WithLumberjackMaxBackups(maxBackups int) Option {
	return option(func(cfg *LumberjackConfig) {
		cfg.maxBackups = maxBackups
	})
}

func WithLumberjackCompress(compress bool) Option {
	return option(func(cfg *LumberjackConfig) {
		cfg.compress = compress
	})
}

func WithLumberjackLocalTime(localTime bool) Option {
	return option(func(cfg *LumberjackConfig) {
		cfg.localTime = localTime
	})
}
