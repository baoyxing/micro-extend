package client

import "time"

// Client **********************************客户端对象******************************
// 客户端配置
type Client struct {
	TimeoutControl timeOutControl `json:"timeout_control" mapstructure:"timeout_control" yaml:"timeout_control"` //超时控制
	ConnectionType connectionType `json:"connection_type" mapstructure:"connection_type" yaml:"connection_type"` // 连接类型
	FailureRetry   failureRetry   `json:"failure_retry" mapstructure:"failure_retry" yaml:"failure_retry"`       //请求重试
	LoadBalancer   loadBalancer   `json:"load_balancer" mapstructure:"load_balancer" yaml:"load_balancer"`       //负载均衡
	CBSuite        cbsuite        `json:"cb_suite" mapstructure:"cb_suite" yaml:"cb_suite"`                      //熔断器
	StatsLevel     StatsLevel     `json:"stats_level" mapstructure:"stats_level" yaml:"stats_level"`             //埋点策略&埋点粒度
}

// Polaris 注册中心配置
type Polaris struct {
	Enable  bool   `json:"enable" mapstructure:"enable" yaml:"enable"` //是否启用注册中心，默认开启
	Network string `json:"network" mapstructure:"network" yaml:"network"`
	Address string `json:"address" mapstructure:"address" yaml:"address"`
}

// Jaeger 链路追踪配置
type Jaeger struct {
	Enable   bool   `json:"enable" mapstructure:"enable" yaml:"enable"`       //是否启用链路追踪
	Endpoint string `json:"endpoint" mapstructure:"endpoint" yaml:"endpoint"` //地址
}

// 超时控制
type timeOutControl struct {
	RpcTimeout     rpcTimeout     `json:"rpc_timeout" mapstructure:"rpc_timeout" yaml:"rpc_timeout"`
	ConnectTimeOut connectTimeOut `json:"connect_time_out" mapstructure:"connect_time_out" yaml:"connect_time_out"`
}

// 连接类型（长链接 短链接）
type connectionType struct {
	ShortConnection shortConnection `json:"short_connection" mapstructure:"short_connection" yaml:"short_connection"` //短链接
	LongConnection  longConnection  `json:"long_connection" mapstructure:"long_connection" yaml:"long_connection"`    //长链接
	ClientTransport clientTransport `json:"client_transport" mapstructure:"client_transport" yaml:"client_transport"` //客户端多路复用

}

// rpc超时控制
type rpcTimeout struct {
	Enable  bool   `json:"enable" mapstructure:"enable" yaml:"enable"`       //是否启用rpc超时
	Timeout string `json:"time_out" mapstructure:"time_out" yaml:"time_out"` //超时时间 （默认 1s 单位："ns", "us" (or "µs"), "ms", "s", "m", "h"）
}

// connect超时控制
type connectTimeOut struct {
	Enable  bool   `json:"enable" mapstructure:"enable" yaml:"enable"`       //是否启用rpc超时
	TimeOut string `json:"time_out" mapstructure:"time_out" yaml:"time_out"` //连接超时 （默认：50ms）
}

// 短链接
type shortConnection struct {
	Enable bool `json:"enable" mapstructure:"enable" yaml:"enable"` //是否启用短链接
}

// 长链接
type longConnection struct {
	Enable            bool   `json:"enable" mapstructure:"enable" yaml:"enable"`                                           //是否启用长链接
	MaxIdlePerAddress int    `json:"max_idle_per_address" mapstructure:"max_idle_per_address" yaml:"max_idle_per_address"` //最大空闲地址
	MinIdlePerAddress int    `json:"min_idle_per_address" mapstructure:"min_idle_per_address" yaml:"min_idle_per_address"` //最小空闲地址
	MaxIdleGlobal     int    `json:"max_idle_global" mapstructure:"max_idle_global" yaml:"max_idle_global"`                //最大空闲数
	MaxIdleTimeOut    string `json:"max_idle_time_out" mapstructure:"max_idle_time_out" yaml:"max_idle_time_out"`          //最大空闲超时
}

// 客户端多路复用
type clientTransport struct {
	Enable        bool `json:"enable" mapstructure:"enable" yaml:"enable"`                         //是否启用多路复用
	MuxConnection int  `json:"mux_connection" mapstructure:"mux_connection" yaml:"mux_connection"` //连接数
}

// 重试机制
type failureRetry struct {
	Enable        bool `json:"enable" mapstructure:"enable" yaml:"enable"`                            //是否启用请求重试机制
	MaxRetryTimes int  `json:"max_retry_times" mapstructure:"max_retry_times" yaml:"max_retry_times"` //重试次数
}

// 负载均衡
type loadBalancer struct {
	Enable bool `json:"enable" mapstructure:"enable" yaml:"enable"` //是否启用负载均衡
}

// 熔断器
type cbsuite struct {
	Enable bool `json:"enable" mapstructure:"enable" yaml:"enable"` //是否启用熔断器
}

// **********************************公共对象*******************************
type StatsLevel struct {
	LevelDisabled bool `json:"level_disabled" mapstructure:"level_disabled" yaml:"level_disabled"`
	LevelBase     bool `json:"level_base" mapstructure:"level_base" yaml:"level_base"`
	LevelDetailed bool `json:"level_detailed" mapstructure:"level_detailed" yaml:"level_detailed"`
}

type RpcClientConf struct {
	Addr             string
	MuxConnectionNum int
	RpcTimeout       time.Duration
	ProviderEndpoint string
	ServiceName      string
}

type RPCServerOption struct {
	Name     string `json:"name" mapstructure:"name" yaml:"name"`
	Intranet string `json:"intranet"  mapstructure:"intranet" yaml:"intranet"`
}
