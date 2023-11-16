package logutils

import (
	Kitexzap "github.com/baoyxing/micro-extend/pkg/utils/logutils/hertz"
	"github.com/cloudwego/kitex/pkg/klog"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

func NewKitexLog(opts ...Option) {
	config := DefaultOptionConfig()
	for _, opt := range opts {
		opt.apply(config)
	}
	dynamicLevel := zap.NewAtomicLevel()
	dynamicLevel.SetLevel(zap.DebugLevel)
	coreConfigs := make([]Kitexzap.CoreConfig, 0)
	switch config.outputMode {
	case 1:
		coreConfigs = append(coreConfigs, Kitexzap.CoreConfig{
			Enc: zapcore.NewConsoleEncoder(pkgEncoderConfig()),
			Ws:  zapcore.AddSync(os.Stdout),
			Lvl: dynamicLevel,
		})
	case 2:
		coreConfigs = append(coreConfigs, Kitexzap.CoreConfig{
			Enc: zapcore.NewJSONEncoder(pkgEncoderConfig()),
			Ws:  NewWriteSyncer("/all/app", config),
			Lvl: zap.NewAtomicLevelAt(zapcore.DebugLevel),
		})
		coreConfigs = append(coreConfigs, Kitexzap.CoreConfig{
			Enc: zapcore.NewJSONEncoder(pkgEncoderConfig()),
			Ws:  NewWriteSyncer("/debug/app", config),
			Lvl: zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
				return lev == zap.DebugLevel
			}),
		})
		coreConfigs = append(coreConfigs, Kitexzap.CoreConfig{
			Enc: zapcore.NewJSONEncoder(pkgEncoderConfig()),
			Ws:  NewWriteSyncer("/info/app", config),
			Lvl: zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
				return lev == zap.InfoLevel
			}),
		})
		coreConfigs = append(coreConfigs, Kitexzap.CoreConfig{
			Enc: zapcore.NewJSONEncoder(pkgEncoderConfig()),
			Ws:  NewWriteSyncer("/warn/app", config),
			Lvl: zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
				return lev == zap.WarnLevel
			}),
		})
		coreConfigs = append(coreConfigs, Kitexzap.CoreConfig{
			Enc: zapcore.NewJSONEncoder(pkgEncoderConfig()),
			Ws:  NewWriteSyncer("/error/app", config),
			Lvl: zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
				return lev == zap.ErrorLevel
			}),
		})
		coreConfigs = append(coreConfigs, Kitexzap.CoreConfig{
			Enc: zapcore.NewJSONEncoder(pkgEncoderConfig()),
			Ws:  NewWriteSyncer("/panic/app", config),
			Lvl: zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
				return lev >= zap.DPanicLevel
			}),
		})
	case 3:
		coreConfigs = append(coreConfigs, Kitexzap.CoreConfig{
			Enc: zapcore.NewConsoleEncoder(pkgEncoderConfig()),
			Ws:  zapcore.AddSync(os.Stdout),
			Lvl: dynamicLevel,
		})
		coreConfigs = append(coreConfigs, Kitexzap.CoreConfig{
			Enc: zapcore.NewJSONEncoder(pkgEncoderConfig()),
			Ws:  NewWriteSyncer("/all/app", config),
			Lvl: zap.NewAtomicLevelAt(zapcore.DebugLevel),
		})
		coreConfigs = append(coreConfigs, Kitexzap.CoreConfig{
			Enc: zapcore.NewJSONEncoder(pkgEncoderConfig()),
			Ws:  NewWriteSyncer("/debug/app", config),
			Lvl: zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
				return lev == zap.DebugLevel
			}),
		})
		coreConfigs = append(coreConfigs, Kitexzap.CoreConfig{
			Enc: zapcore.NewJSONEncoder(pkgEncoderConfig()),
			Ws:  NewWriteSyncer("/info/app", config),
			Lvl: zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
				return lev == zap.InfoLevel
			}),
		})
		coreConfigs = append(coreConfigs, Kitexzap.CoreConfig{
			Enc: zapcore.NewJSONEncoder(pkgEncoderConfig()),
			Ws:  NewWriteSyncer("/warn/app", config),
			Lvl: zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
				return lev == zap.WarnLevel
			}),
		})
		coreConfigs = append(coreConfigs, Kitexzap.CoreConfig{
			Enc: zapcore.NewJSONEncoder(pkgEncoderConfig()),
			Ws:  NewWriteSyncer("/error/app", config),
			Lvl: zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
				return lev == zap.ErrorLevel
			}),
		})
		coreConfigs = append(coreConfigs, Kitexzap.CoreConfig{
			Enc: zapcore.NewJSONEncoder(pkgEncoderConfig()),
			Ws:  NewWriteSyncer("/panic/app", config),
			Lvl: zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
				return lev >= zap.DPanicLevel
			}),
		})
	}
	logger := Kitexzap.NewLogger(Kitexzap.WithCores(coreConfigs...),
		Kitexzap.WithZapOptions(zap.AddCaller(), zap.AddCallerSkip(3)),
	)

	defer func(logger *Kitexzap.Logger) {
		err := logger.Sync()
		if err != nil {
			logger.Error(err)
		}
	}(logger)
	klog.SetLogger(logger)
}
