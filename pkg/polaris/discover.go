package polaris

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/kitex-contrib/polaris"
)

//NewDiscover 服务发现
func NewDiscover(ctx context.Context, nameSpace string,
	log hlog.CtxLogger) *polaris.ClientSuite {
	resolver, err := polaris.NewPolarisResolver(polaris.ClientOptions{})
	if err != nil {
		log.CtxErrorf(ctx, "NewPolarisResolver creates a polaris based resolver:%v error：%v", resolver, err)
		return nil
	}
	balancer, err := polaris.NewPolarisBalancer()
	if err != nil {
		log.CtxErrorf(ctx, "NewPolarisBalancer creates a polaris based balancer:%v error：%v", balancer, err)
		return nil
	}
	return &polaris.ClientSuite{
		DstNameSpace:       nameSpace,
		Resolver:           resolver,
		Balancer:           balancer,
		ReportCallResultMW: polaris.NewUpdateServiceCallResultMW(),
	}
}
