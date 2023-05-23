package log

// Logger 日志配置
type Logger struct {
	Enable     bool   `json:"enable" mapstructure:"enable" yaml:"enable"`                //是否启用自定义日志配置
	Filename   string `json:"filename" mapstructure:"filename" yaml:"filename"`          //路径
	MaxSize    int    `json:"max_size" mapstructure:"max_size" yaml:"max_size"`          //日志的最大大小（M）
	MaxBackups int    `json:"max_backups" mapstructure:"max_backups" yaml:"max_backups"` //日志的最大保存数量
	MaxAge     int    `json:"max_age" mapstructure:"max_age" yaml:"max_age"`             //日志文件存储最大天数
	Compress   bool   `json:"compress" mapstructure:"compress" yaml:"compress"`          //是否执行压缩
	LocalTime  bool   `json:"local_time" mapstructure:"local_time" yaml:"local_time"`    //是否使用格式化时间辍
	Level      string `json:"level" mapstructure:"level" yaml:"level"`                   // 日志等级  【trace,debug,info,notice,warn,error,fatal】
}
