package rpc

import (
	"context"
	"github.com/baoyxing/micro-extend/pkg/config/kitex_conf"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/registry"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/stats"
	"github.com/cloudwego/kitex/pkg/transmeta"
	"github.com/cloudwego/kitex/server"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	"github.com/kitex-contrib/polaris"

	"net"
)

func ServerOptions(confServer kitex_conf.Server,
	confService kitex_conf.Service, log klog.CtxLogger) ([]server.Option, error) {
	var options []server.Option
	ctx := context.Background()
	if confServer.Rpc.Enable {
		addr, err := net.ResolveTCPAddr(confServer.Rpc.Network, confServer.Rpc.Address)
		if err != nil {
			log.CtxErrorf(ctx, "ResolveTCPAddr error addr:%v err:%v", addr, err)
			return nil, err
		}
		options = append(options, server.WithServiceAddr(addr))
		log.CtxInfof(ctx, "服务端配置自定义端口已配置成功 %v", addr)
	}

	if confServer.Polaris.Enable {
		r, err := polaris.NewPolarisRegistry(polaris.ServerOptions{})
		if err != nil {
			log.CtxErrorf(ctx, "polaris NewPolarisRegistry fatal ：%v", err)
			return nil, err
		}
		info := &registry.Info{
			ServiceName: confService.Name,
			Tags: map[string]string{
				polaris.NameSpaceTagKey: confService.Space,
			},
		}
		options = append(options, server.WithRegistry(r))
		options = append(options, server.WithRegistryInfo(info))
		log.CtxInfof(ctx, "服务端配置北极星注册中心已配置成功 name：%v，nameSpace:%v", confService.Name, confService.Space)
	}

	//是否启用jaeger链路追踪
	if confServer.Jaeger.Enable {
		provider.NewOpenTelemetryProvider(
			provider.WithServiceName(confService.Name),
			provider.WithExportEndpoint(confServer.Jaeger.Endpoint),
			provider.WithInsecure(),
		)
		options = append(options, server.WithSuite(tracing.NewServerSuite()))
		options = append(options, server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: confService.Space}))
		log.CtxInfof(ctx, "服务端配置链路已配置成功 ServiceName：%s，Endpoint:%s", confService.Space, confServer.Jaeger.Endpoint)
	}
	//是否启用多路复用
	if confServer.Transport.Enable {
		options = append(options, server.WithMuxTransport())
		log.CtxInfof(ctx, "服务端配置多路复用已配置成功 ServiceName：%v", confService.Name)
	}
	//是否启用限流器
	if confServer.Limit.Enable {
		options = append(options, server.WithLimit(&limit.Option{
			MaxConnections: confServer.Limit.MaxConnections,
			MaxQPS:         confServer.Limit.MaxQPS,
		}))
		log.CtxInfof(ctx, "服务端配置限流器已配置成功 ServiceName：%v 最大连接数:%d 最大qps:%d",
			confService.Name, confServer.Limit.MaxConnections, confServer.Limit.MaxQPS)
	}

	//埋点策略&埋点粒度
	if confServer.StatsLevel.LevelBase {
		options = append(options, server.WithStatsLevel(stats.LevelBase))
		log.CtxInfof(ctx, "客户端配置启用基本埋点 已启用 LevelBase：%v", stats.LevelBase)
	}
	if confServer.StatsLevel.LevelDetailed {
		options = append(options, server.WithStatsLevel(stats.LevelDetailed))
		log.CtxInfof(ctx, "客户端配置启用基本埋点和细粒度埋点 已启用 LevelDetailed：%v", stats.LevelDetailed)
	}
	if confServer.StatsLevel.LevelDisabled {
		options = append(options, server.WithStatsLevel(stats.LevelDisabled))
		log.CtxInfof(ctx, "客户端配置禁用埋点 已禁用 LevelDisabled：%v", stats.LevelDisabled)
	}
	options = append(options, server.WithMetaHandler(transmeta.ServerTTHeaderHandler))
	return options, nil
}
