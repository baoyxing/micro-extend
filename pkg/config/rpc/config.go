package rpc

type ClientOption struct {
	Addr             string
	MuxConnectionNum int64
	RpcTimeout       int64
	ProviderEndpoint string
}
