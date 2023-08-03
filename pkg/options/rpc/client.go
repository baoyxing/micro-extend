package rpc

import (
	"context"
	clientConf "github.com/baoyxing/micro-extend/pkg/config/client"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/connpool"
	"github.com/cloudwego/kitex/pkg/loadbalance"
	"github.com/cloudwego/kitex/pkg/retry"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/stats"
	"github.com/cloudwego/kitex/pkg/transmeta"
	"github.com/cloudwego/kitex/transport"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	"github.com/kitex-contrib/polaris"
	"time"
)

type ctxKey int

const (
	ctxConsistedKey ctxKey = iota
)

func ClientOptions(confClient clientConf.Client, polaris clientConf.Polaris, jaeger clientConf.Jaeger,
	clientName string, suite *polaris.ClientSuite,
	log hlog.CtxLogger) ([]client.Option, error) {
	ctx := context.Background()
	var options []client.Option
	if len(confClient.TimeoutControl.RpcTimeout.Timeout) > 1 && confClient.TimeoutControl.RpcTimeout.Enable {
		duration, err := time.ParseDuration(confClient.TimeoutControl.RpcTimeout.Timeout)
		if err != nil {
			log.CtxErrorf(context.Background(), "ParseDuration RpcTimeout duration:%v error：%v", duration, err)
			return nil, err
		}
		options = append(options, client.WithRPCTimeout(duration))
		log.CtxInfof(ctx, "客户端配置RPC超时配置已配置成功 %v", duration)
	} else {
		// 未配置超时，则默认 1s
		options = append(options, client.WithRPCTimeout(time.Second*1))
		log.CtxInfof(ctx, "客户端配置RPC超时配置已启用默认配置 %v", time.Second*1)
	}

	if len(confClient.TimeoutControl.ConnectTimeOut.TimeOut) > 1 && confClient.TimeoutControl.ConnectTimeOut.Enable {
		duration, err := time.ParseDuration(confClient.TimeoutControl.ConnectTimeOut.TimeOut)
		if err != nil {
			log.CtxErrorf(ctx, "ParseDuration ConnectTimeOut duration:%v error：%v", duration, err)
			return nil, err
		}
		options = append(options, client.WithConnectTimeout(duration))
		log.CtxInfof(ctx, "客户端配置连接超时已配置成功 %v", duration)
	} else {
		// 未配置超时，则默认 1s
		options = append(options, client.WithConnectTimeout(time.Millisecond*50))
		log.CtxInfof(ctx, "客户端配置连接超时配置已启用默认配置 %v", time.Millisecond*50)
	}

	//客户端配置jaeger
	if jaeger.Enable {
		provider.NewOpenTelemetryProvider(
			provider.WithServiceName(clientName),
			provider.WithExportEndpoint(jaeger.Endpoint),
			provider.WithEnableTracing(true),
			provider.WithInsecure(),
		)
		//defer p.Shutdown(ctx)
		options = append(options, client.WithSuite(tracing.NewClientSuite()))
		options = append(options, client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: clientName}))
		log.CtxInfof(ctx, "客户端配置链路已配置成功 ClientName：%v，Endpoint:%v", clientName, jaeger.Endpoint)
	}

	// 客户端连接服务中心
	if polaris.Enable {
		options = append(options, client.WithSuite(suite))
		log.CtxInfof(ctx, "客户端配置连接配置中心已配置成功 ClientName：%v，Endpoint:%v", clientName, polaris.Address)
	}

	//客户端连接类型---短链接  3选1(短连接、长链接、多路复用)
	if confClient.ConnectionType.ShortConnection.Enable {

		options = append(options, client.WithShortConnection())
		log.CtxInfof(ctx, "客户端配置短链接配置已启用 %v", confClient.ConnectionType.ShortConnection.Enable)
	}
	//客户端连接类型---长链接 3选1(短连接、长链接、多路复用)
	if confClient.ConnectionType.LongConnection.Enable {
		duration, err := time.ParseDuration(confClient.ConnectionType.LongConnection.MaxIdleTimeOut)
		if err != nil {
			log.CtxErrorf(ctx, "ParseDuration LongConnection MaxIdleTimeOut duration:%v error：%v", duration, err)
			return nil, err
		}
		pool := connpool.IdleConfig{
			MinIdlePerAddress: confClient.ConnectionType.LongConnection.MinIdlePerAddress,
			MaxIdlePerAddress: confClient.ConnectionType.LongConnection.MaxIdlePerAddress,
			MaxIdleGlobal:     confClient.ConnectionType.LongConnection.MaxIdleGlobal,
			MaxIdleTimeout:    duration,
		}
		options = append(options, client.WithLongConnection(pool))
		log.CtxInfof(ctx, "客户端配置长链接配置已启用 连接池：%v", pool)

	}

	//客户端连接类型--多路复用 3选1(短连接、长链接、多路复用)
	if confClient.ConnectionType.ClientTransport.Enable {
		options = append(options, client.WithMuxConnection(confClient.ConnectionType.ClientTransport.MuxConnection))
		log.CtxInfof(ctx, "客户端配置多路复用 已启用 ：%v", confClient.ConnectionType.ClientTransport.MuxConnection)
	}

	//请求重试机制
	if confClient.FailureRetry.Enable {
		failurePolicy := retry.NewFailurePolicy()
		failurePolicy.WithMaxRetryTimes(confClient.FailureRetry.MaxRetryTimes)
		options = append(options, client.WithFailureRetry(failurePolicy))
		log.CtxInfof(ctx, "客户端配置请求重试机制 已启用 重试次数：%v", confClient.FailureRetry.MaxRetryTimes)
	}

	//负载均衡
	if confClient.LoadBalancer.Enable {
		options = append(options, client.WithLoadBalancer(loadbalance.NewConsistBalancer(
			loadbalance.NewConsistentHashOption(func(ctx context.Context, request interface{}) string {
				return ctx.Value(ctxConsistedKey).(string)
			}))))
		log.CtxInfof(ctx, "客户端配置负载均衡 已启用 ：%v", confClient.LoadBalancer.Enable)
	}
	//熔断器
	if confClient.CBSuite.Enable {
		//options = append(options, client.WithCircuitBreaker(circuitbreak.NewCBSuite(func(ri rpcinfo.RPCInfo) string {
		//	return ""
		//})))
	}
	//埋点策略&埋点粒度
	if confClient.StatsLevel.LevelBase {
		options = append(options, client.WithStatsLevel(stats.LevelBase))
		log.CtxInfof(ctx, "客户端配置启用基本埋点 已启用 LevelBase：%v", stats.LevelBase)
	}
	if confClient.StatsLevel.LevelDetailed {
		options = append(options, client.WithStatsLevel(stats.LevelDetailed))
		log.CtxInfof(ctx, "客户端配置启用基本埋点和细粒度埋点 已启用 LevelDetailed：%v", stats.LevelDetailed)
	}
	if confClient.StatsLevel.LevelDisabled {
		options = append(options, client.WithStatsLevel(stats.LevelDisabled))
		log.CtxInfof(ctx, "客户端配置禁用埋点 已禁用 LevelDisabled：%v", stats.LevelDisabled)
	}

	options = append(options, client.WithMetaHandler(transmeta.ClientTTHeaderHandler))
	options = append(options, client.WithTransportProtocol(transport.TTHeader))

	return options, nil
}
