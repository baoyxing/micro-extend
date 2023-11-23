package http

import (
	"context"
	"github.com/baoyxing/micro-extend/pkg/config/hertz_conf"
	serverconfig "github.com/cloudwego/hertz/pkg/common/config"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/hertz-contrib/obs-opentelemetry/provider"
	hertztracing "github.com/kitex-contrib/obs-opentelemetry/tracing"
)

func NewServerTracer(server hertz_conf.Server, service hertz_conf.Service,
	log hlog.CtxLogger) (serverconfig.Option, *hertztracing.Config) {
	provider.NewOpenTelemetryProvider(
		provider.WithServiceName(service.Name),
		provider.WithExportEndpoint(server.Jaeger.Endpoint),
		provider.WithEnableTracing(true),
		provider.WithInsecure(),
	)
	log.CtxInfof(context.Background(),
		"服务端配置链路已配置成功 ServiceName：%v，Endpoint:%v",
		service.Name, server.Jaeger.Endpoint)
	return hertztracing.NewServerTracer()
}
