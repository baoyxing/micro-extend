package tools

import (
	"context"
	rsp "github.com/baoyxing/hertz-contrib/pkg/app"
	"github.com/bytedance/sonic"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/pkg/errors"
	"strconv"
)

func ParseHttpBody(ctx *app.RequestContext, teaKey string) error {
	body, err := ctx.Body()
	if err != nil {
		err = errors.Wrap(err, "body not null")
		return err
	}
	tempBodyLengthStr := ctx.Request.Header.Get("X-Body-Length")
	if tempBodyLengthStr == "" {
		return errors.New("X-Body-Length not null")
	}
	bodyLength, err := strconv.Atoi(tempBodyLengthStr)
	if err != nil {
		err = errors.Wrap(err, "X-Body-Length not number")
		return err
	}
	body, err = TeaHexDecode(body, bodyLength, teaKey)
	if err != nil {
		err = errors.Wrap(err, "body tea decode failure")
		return err
	}
	ctx.Request.SetBody(body)
	return nil
}
func EncodeHttpBody(c context.Context, ctx *app.RequestContext, data interface{}, teaKey string) {
	src, err := sonic.Marshal(data)
	if err != nil {
		hlog.CtxErrorf(c, "respon marshal failure,err:%s", err.Error())
		ctx.AbortWithMsg("非法Body", 403)
		return
	}
	dataByte, _, err := EncodeTeaStr(src, teaKey)
	if err != nil {
		hlog.CtxErrorf(c, "Encode tea failure,err:%s", err.Error())
		ctx.AbortWithMsg("非法Body", 403)
		return
	}
	rsp.OkWithBody(dataByte, ctx, len(src))
}
