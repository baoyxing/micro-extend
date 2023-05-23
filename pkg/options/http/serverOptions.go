package http

import (
	"context"
	"github.com/baoyxing/micro-extend/pkg/config/hertz_conf"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/app/server/registry"
	"github.com/cloudwego/hertz/pkg/common/config"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/hertz-contrib/registry/polaris"
	"time"
)

func ServerOptions(hServer hertz_conf.Server,
	service hertz_conf.Service, log hlog.CtxLogger, option config.Option) []config.Option {

	var options []config.Option

	ctx := context.Background()

	////是否启用jaeger链路追踪
	if hServer.Jaeger.Enable {
		options = append(options, option)
	}

	// 自定义 http
	if hServer.Http.Enable {
		options = append(options, server.WithHostPorts(hServer.Http.Address))
		options = append(options, server.WithExitWaitTime(time.Duration(hServer.Http.ExitWaitTime)*time.Second))
		log.CtxInfof(ctx, "服务端配置自定义端口已配置成功 %v", hServer.Http.Address)
		log.CtxInfof(ctx, "服务端配置优雅退出配置成功 %v s", hServer.Http.ExitWaitTime)
	}

	//是否启用北极星注册中心
	if hServer.Polaris.Enable {
		r, err := polaris.NewPolarisRegistry()
		if err != nil {
			log.CtxErrorf(ctx, "polaris NewPolarisRegistry fatal ：%v", err)
			return nil
		}
		info := &registry.Info{
			ServiceName: service.Name,
			Addr:        utils.NewNetAddr(hServer.Polaris.Network, hServer.Polaris.Address),
			Tags: map[string]string{
				"namespace": service.Space,
			},
		}
		options = append(options, server.WithRegistry(r, info))
		log.CtxInfof(ctx, "服务端配置北极星注册中心已配置成功 Network: %s,Address:%s,name：%v，nameSpace:%v",
			hServer.Polaris.Network, hServer.Polaris.Address, service.Name, service.Space)
	}

	return options
}
