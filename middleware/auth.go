package middleware

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol"
	"github.com/pkg/errors"
	"strings"
)

func Auth(sk string) app.HandlerFunc {
	return func(context context.Context, ctx *app.RequestContext) {
		ok, err, kvs := authRequest(ctx.Request)
		if err != nil {
			ctx.AbortWithStatus(400)
			return
		}
		if !ok {
			ctx.AbortWithStatus(403)
			return
		}
		ak := kvs["ak"]
		t := kvs["t"]
		url := ctx.Request.URI()
		signed := customerSigned(string(url.Path()), t, ak, sk)
		if signed != kvs["sign"] {
			ctx.AbortWithStatus(403)
			return
		}
		ctx.Next(context)
	}
}

func authRequest(req protocol.Request) (bool, error, map[string]string) {
	authHeader := req.Header.Get("Authorization")
	if authHeader == "" {
		return false, errors.New("authorization empty"), nil
	}
	authHeaderBytes, err := base64.StdEncoding.DecodeString(authHeader)
	if err != nil {
		return false, err, nil
	}
	kvs := make(map[string]string)
	for _, v := range strings.Split(string(authHeaderBytes), "&") {
		if strings.Contains(v, "=") {
			kv := strings.Split(v, "=")
			if len(kv) == 2 {
				kvs[kv[0]] = kv[1]
			}
		}
	}
	if _, ok := kvs["ak"]; !ok || kvs["ak"] == "" {
		return false, errors.New("ak empty"), nil
	}

	if _, ok := kvs["sign"]; !ok || kvs["sign"] == "" {
		return false, errors.New("sign empty"), nil
	}

	if _, ok := kvs["t"]; !ok || kvs["t"] == "" {
		return false, errors.New("t empty"), nil
	}
	return true, nil, kvs
}

func customerSigned(path string, t string, ak string, sk string) string {
	msg := fmt.Sprintf("%s%s/%s", ak, path, t)
	h := hmac.New(sha256.New, []byte(sk))
	h.Write([]byte(msg))
	sign := h.Sum(nil)
	return fmt.Sprintf("%x", sign)
}
