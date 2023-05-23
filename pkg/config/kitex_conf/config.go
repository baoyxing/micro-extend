package kitex_conf

import "github.com/baoyxing/micro-extend/pkg/config/client"

// Server 服务端配置
type Server struct {
	Rpc        rpc               `json:"rpc" mapstructure:"rpc" yaml:"rpc"`                         //服务ip配置
	Polaris    client.Polaris    `json:"polaris" mapstructure:"polaris" yaml:"polaris"`             //北极星注册中心配置
	Jaeger     client.Jaeger     `json:"jaeger" mapstructure:"jaeger" yaml:"jaeger"`                //链路配置
	Transport  transport         `json:"transport" mapstructure:"transport" yaml:"transport"`       //多路复用配置
	Limit      limit             `json:"limit" mapstructure:"limit" yaml:"limit"`                   //限流器
	StatsLevel client.StatsLevel `json:"stats_level" mapstructure:"stats_level" yaml:"stats_level"` //埋点策略&埋点粒度
}

// Service 服务名称配置
type Service struct {
	Space   string `json:"space" mapstructure:"space" yaml:"space"`       //服务空间名称
	Name    string `json:"name" mapstructure:"name" yaml:"name"`          //服务名称
	Version string `json:"version" mapstructure:"version" yaml:"version"` //版本信息
}

// 服务地址端口配置
type rpc struct {
	Enable  bool   `json:"enable" mapstructure:"enable" yaml:"enable" `   //是否启用rpc自定义配置
	Address string `json:"address" mapstructure:"address" yaml:"address"` //地址
	Network string `json:"network" mapstructure:"network" yaml:"network"` //连接方式 (tcp udp)
}

// 多路复用配置
type transport struct {
	Enable bool `json:"enable" mapstructure:"enable" yaml:"enable"` //是否启用多路复用
}

// 限流器配置
type limit struct {
	Enable         bool `json:"enable" mapstructure:"enable" yaml:"enable"`                            //是否启用多路复用
	MaxConnections int  `json:"max_connections" mapstructure:"max_connections" yaml:"max_connections"` // 最大连接数
	MaxQPS         int  `json:"max_qps" mapstructure:"max_qps" yaml:"max_qps"`                         //最大qps
}
