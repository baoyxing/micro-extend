package logutils

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	hertzzap "github.com/hertz-contrib/logger/zap"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

func NewHertzLog(opts ...Option) {
	config := DefaultOptionConfig()
	for _, opt := range opts {
		opt.apply(config)
	}
	dynamicLevel := zap.NewAtomicLevel()
	dynamicLevel.SetLevel(zap.DebugLevel)

	coreConfigs := make([]hertzzap.CoreConfig, 0)
	switch config.outputMode {
	case 1:
		coreConfigs = append(coreConfigs, hertzzap.CoreConfig{
			Enc: zapcore.NewConsoleEncoder(pkgEncoderConfig()),
			Ws:  zapcore.AddSync(os.Stdout),
			Lvl: dynamicLevel,
		})
	case 2:
		coreConfigs = append(coreConfigs, hertzzap.CoreConfig{
			Enc: zapcore.NewJSONEncoder(pkgEncoderConfig()),
			Ws:  NewWriteSyncer("/all/app", config),
			Lvl: zap.NewAtomicLevelAt(zapcore.DebugLevel),
		})
		coreConfigs = append(coreConfigs, hertzzap.CoreConfig{
			Enc: zapcore.NewJSONEncoder(pkgEncoderConfig()),
			Ws:  NewWriteSyncer("/debug/app", config),
			Lvl: zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
				return lev == zap.DebugLevel
			}),
		})
		coreConfigs = append(coreConfigs, hertzzap.CoreConfig{
			Enc: zapcore.NewJSONEncoder(pkgEncoderConfig()),
			Ws:  NewWriteSyncer("/info/app", config),
			Lvl: zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
				return lev == zap.InfoLevel
			}),
		})
		coreConfigs = append(coreConfigs, hertzzap.CoreConfig{
			Enc: zapcore.NewJSONEncoder(pkgEncoderConfig()),
			Ws:  NewWriteSyncer("/warn/app", config),
			Lvl: zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
				return lev == zap.WarnLevel
			}),
		})
		coreConfigs = append(coreConfigs, hertzzap.CoreConfig{
			Enc: zapcore.NewJSONEncoder(pkgEncoderConfig()),
			Ws:  NewWriteSyncer("/error/app", config),
			Lvl: zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
				return lev == zap.ErrorLevel
			}),
		})
		coreConfigs = append(coreConfigs, hertzzap.CoreConfig{
			Enc: zapcore.NewJSONEncoder(pkgEncoderConfig()),
			Ws:  NewWriteSyncer("/panic/app", config),
			Lvl: zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
				return lev >= zap.DPanicLevel
			}),
		})
	case 3:
		coreConfigs = append(coreConfigs, hertzzap.CoreConfig{
			Enc: zapcore.NewConsoleEncoder(pkgEncoderConfig()),
			Ws:  zapcore.AddSync(os.Stdout),
			Lvl: dynamicLevel,
		})
		coreConfigs = append(coreConfigs, hertzzap.CoreConfig{
			Enc: zapcore.NewJSONEncoder(pkgEncoderConfig()),
			Ws:  NewWriteSyncer("/all/app", config),
			Lvl: zap.NewAtomicLevelAt(zapcore.DebugLevel),
		})
		coreConfigs = append(coreConfigs, hertzzap.CoreConfig{
			Enc: zapcore.NewJSONEncoder(pkgEncoderConfig()),
			Ws:  NewWriteSyncer("/debug/app", config),
			Lvl: zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
				return lev == zap.DebugLevel
			}),
		})
		coreConfigs = append(coreConfigs, hertzzap.CoreConfig{
			Enc: zapcore.NewJSONEncoder(pkgEncoderConfig()),
			Ws:  NewWriteSyncer("/info/app", config),
			Lvl: zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
				return lev == zap.InfoLevel
			}),
		})
		coreConfigs = append(coreConfigs, hertzzap.CoreConfig{
			Enc: zapcore.NewJSONEncoder(pkgEncoderConfig()),
			Ws:  NewWriteSyncer("/warn/app", config),
			Lvl: zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
				return lev == zap.WarnLevel
			}),
		})
		coreConfigs = append(coreConfigs, hertzzap.CoreConfig{
			Enc: zapcore.NewJSONEncoder(pkgEncoderConfig()),
			Ws:  NewWriteSyncer("/error/app", config),
			Lvl: zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
				return lev == zap.ErrorLevel
			}),
		})
		coreConfigs = append(coreConfigs, hertzzap.CoreConfig{
			Enc: zapcore.NewJSONEncoder(pkgEncoderConfig()),
			Ws:  NewWriteSyncer("/panic/app", config),
			Lvl: zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
				return lev >= zap.DPanicLevel
			}),
		})
	}
	logger := hertzzap.NewLogger(hertzzap.WithCores(coreConfigs...),
		hertzzap.WithZapOptions(zap.AddCaller(), zap.AddCallerSkip(3)), // 行号
	)
	defer logger.Sync()
	hlog.SetLogger(logger)
}
