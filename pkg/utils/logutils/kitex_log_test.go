package logutils

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"testing"
	"time"
)

func TestNewKitexLog(t *testing.T) {
	opts := make([]Option, 0, 8)
	opts = append(opts, WithPath("./log"))
	opts = append(opts, WithMaxSize(10))
	opts = append(opts, WithMaxBackups(10))
	opts = append(opts, WithMaxAge(30))
	opts = append(opts, WithCompress(false))
	opts = append(opts, WithOutputMode(2))
	opts = append(opts, WithRotationDuration(time.Duration(1)))
	opts = append(opts, WithSuffix(".log"))
	NewKitexLog(opts...)
	klog.SetLevel(klog.LevelDebug)
	log := klog.DefaultLogger()
	log.CtxInfof(context.Background(), "test:%s", "failure")
}
