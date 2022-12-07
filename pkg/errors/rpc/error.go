package rpc

import "github.com/cloudwego/kitex/pkg/kerrors"

type ErrorType int

const (
	ErrorTypeInvalid ErrorType = iota
	// 数据重复错误
	ErrorTypeDBDataRepeat
	//数据库操作错误
	ErrorTypeDBHandle
	// 数据无效
	ErrorTypeDataInvalid
	//network 内部网络处理错误
	ErrorTypeNetwork
	//JsonMarshal json 转换错误
	ErrorTypeJsonMarshal
	// UnMarshal json 解析错误
	ErrorTypeJsonUnMarshal
)

func NewBizStatusError(code ErrorType, err error) kerrors.BizStatusErrorIface {
	return kerrors.NewBizStatusError(int32(code), err.Error())
}

func ParseBizStatusError(err error) (publicErrMsg string, privateErrMsg string) {
	publicErrMsg = "内部错误"
	if bizErr, ok := kerrors.FromBizStatusError(err); ok {
		code := ErrorType(bizErr.BizStatusCode())
		if code > ErrorTypeDBDataRepeat {
			switch code {
			case ErrorTypeDBHandle:
				privateErrMsg += "db handle Error,reason:"
			case ErrorTypeDataInvalid:
				privateErrMsg += "DataInvalid Error,reason:"
			case ErrorTypeNetwork:
				privateErrMsg += "Network Error,reason:"
			case ErrorTypeJsonMarshal:
				privateErrMsg += "JsonMarshal Error,reason:"
			case ErrorTypeJsonUnMarshal:
				privateErrMsg += "JsonUnMarshal Error,reason:"
			}
		}
		privateErrMsg += bizErr.BizMessage()
	} else {
		privateErrMsg += err.Error()
	}

	return
}
