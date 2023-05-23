package hertz_conf

import "github.com/baoyxing/micro-extend/pkg/config/client"

// Server 服务端配置
type Server struct {
	Http      http           `json:"http" mapstructure:"http" yaml:"http"`                //服务ip配置
	Polaris   client.Polaris `json:"polaris" mapstructure:"polaris" yaml:"polaris"`       //北极星注册中心配置
	Auth      auth           `json:"auth" mapstructure:"auth" yaml:"auth"`                // auth 身份认证配置
	Cors      cors           `json:"cors" mapstructure:"cors" yaml:"cors"`                //cors 配置
	Recovery  recovery       `json:"recovery" mapstructure:"recovery" yaml:"recovery"`    // recovery 配置
	Gzip      gzip           `json:"gzip" mapstructure:"gzip" yaml:"gzip"`                // 压缩配置
	I18n      i18n           `json:"i18n" mapstructure:"i18n" yaml:"i18n"`                // 国际化配置
	Swag      swag           `json:"swag" mapstructure:"swag" yaml:"swag"`                // swag文档
	Jaeger    client.Jaeger  `json:"jaeger" mapstructure:"jaeger" yaml:"jaeger"`          //链路配置
	Transport transport      `json:"transport" mapstructure:"transport" yaml:"transport"` //多路复用配置
}

// Service 服务名称配置
type Service struct {
	Space   string `json:"space" mapstructure:"space" yaml:"space"`       //服务空间名称
	Name    string `json:"name" mapstructure:"name" yaml:"name"`          //服务名称
	Version string `json:"version" mapstructure:"version" yaml:"version"` //版本信息
}

// 服务地址端口配置
type http struct {
	Enable       bool   `json:"enable" mapstructure:"enable" yaml:"enable"`                         //是否启用http自定义配置
	Address      string `json:"address" mapstructure:"address" yaml:"address"`                      //地址
	ExitWaitTime int    `json:"exit_wait_time" mapstructure:"exit_wait_time" yaml:"exit_wait_time"` // 配置
}

// 多路复用配置
type transport struct {
	Enable bool `json:"enable" mapstructure:"enable" yaml:"enable"` //是否启用多路复用
}

// auth 配置
type auth struct {
	Enable bool   `json:"enable" mapstructure:"enable" yaml:"enable"` //是否启用auth配置
	AK     string `json:"ak" mapstructure:"ak" yaml:"ak"`
	SK     string `json:"sk" mapstructure:"sk" yaml:"sk"`
	TeaKey string ` json:"tea_key" mapstructure:"tea_key" yaml:"tea_key"`
}

// cors配置
type cors struct {
	Enable bool `json:"enable" mapstructure:"enable" yaml:"enable"` //是否启用cors配置
}

// recovery 配置
type recovery struct {
	Enable bool `json:"enable" mapstructure:"enable" yaml:"enable"` //是否启用 recovery 配置
}

// gzip 压缩配置
type gzip struct {
	Enable             bool     `json:"enable" mapstructure:"enable" yaml:"enable"`                                        //是否启用 gzip 配置
	BestCompression    bool     `json:"best_compression" mapstructure:"best_compression" yaml:"best_compression"`          //提供最佳的文件压缩率
	BestSpeed          bool     `json:"best_speed" mapstructure:"best_speed" yaml:"best_speed"`                            //提供了最佳的压缩速度
	DefaultCompression bool     `json:"default_compression" mapstructure:"default_compression" yaml:"default_compression"` //默认压缩率
	NoCompression      bool     `json:"no_compression" mapstructure:"no_compression" yaml:"no_compression"`                //不进行压缩
	Excluded           excluded `json:"excluded" mapstructure:"excluded" yaml:"excluded"`                                  // 设置不需要压缩的方式
}

type excluded struct {
	Enable              bool                `json:"enable" mapstructure:"enable" yaml:"enable"` //是否启用 excluded 配置
	ExcludedExtensions  excludedExtensions  `json:"excluded_extensions" mapstructure:"excluded_extensions" yaml:"excluded_extensions"`
	ExcludedPaths       excludedPaths       `json:"excluded_paths" mapstructure:"excluded_paths" yaml:"excluded_paths"`
	ExcludedPathRegexes excludedPathRegexes ` json:"excluded_path_regexes" mapstructure:"excluded_path_regexes" yaml:"excluded_path_regexes"`
}

// 设置不需要 gzip 压缩的文件后缀
type excludedExtensions struct {
	Enable     bool   `json:"enable" mapstructure:"enable" yaml:"enable"`             //是否启用 excluded 配置
	Extensions string `json:"extensions" mapstructure:"extensions" yaml:"extensions"` // 文件后缀 数组用,连接 eg:".pdf", ".mp4"
}

// 设置不需要进行 gzip 压缩的文件路径
type excludedPaths struct {
	Enable bool   `json:"enable" mapstructure:"enable" yaml:"enable"` //是否启用 excludedPaths 配置
	Paths  string `json:"paths" mapstructure:"paths" yaml:"paths"`    //文件路径 数组用,连接 eg:/api/
}

// 设置自定义的正则表达式来过滤掉不需要 gzip 压缩的文件
type excludedPathRegexes struct {
	Enable  bool   `json:"enable" mapstructure:"enable" yaml:"enable"`    //是否启用 excludedPathRegexes 配置
	Regexes string `json:"regexes" mapstructure:"regexes" yaml:"regexes"` //文件路径 数组用,连接  eg: /api.*
}

// 设置国际化
type i18n struct {
	Enable bool `json:"enable" mapstructure:"enable" yaml:"enable"` //是否启用 i18n 配置
}

// 设置 swag
type swag struct {
	Enable bool `json:"enable" mapstructure:"enable" yaml:"enable"` //是否启用 swag 配置
}
