package app

import (
	"encoding/hex"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"strconv"
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

const (
	RELOGIN = 1
	ERROR   = 7
	SUCCESS = 0
)

func Result(code int, data interface{}, msg string, ctx *app.RequestContext) {
	if data == nil {
		ctx.JSON(consts.StatusOK, utils.H{
			"code": code,
			"msg":  msg,
		})
	} else {
		ctx.JSON(consts.StatusOK, Response{
			code,
			msg,
			data,
		})
	}
	ctx.Abort()
}

func Ok(ctx *app.RequestContext) {
	Result(SUCCESS, nil, "操作成功", ctx)
}

func OkWithMessage(msg string, ctx *app.RequestContext) {
	Result(SUCCESS, nil, msg, ctx)
}
func OkWithData(data interface{}, ctx *app.RequestContext) {
	Result(SUCCESS, data, "操作成功", ctx)
}

func OkWithDetailed(data interface{}, message string, ctx *app.RequestContext) {
	Result(SUCCESS, data, message, ctx)
}

func Fail(ctx *app.RequestContext) {
	Result(ERROR, nil, "操作失败", ctx)
}

func FailWithMessage(message string, ctx *app.RequestContext) {
	Result(ERROR, nil, message, ctx)
}

func ReLoginWithMessage(message string, ctx *app.RequestContext) {
	Result(RELOGIN, nil, message, ctx)
}

func OkWithBody(data []byte, ctx *app.RequestContext, bodyLength int) {
	dst := make([]byte, hex.EncodedLen(len(data)))
	hex.Encode(dst, data)
	ctx.Header("X-Body-Length", strconv.Itoa(bodyLength))
	ctx.Data(consts.StatusOK, "application/json; charset=UTF-8", dst)
}
