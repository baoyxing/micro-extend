package log

// Logger 日志配置
type Logger struct {
	Enable           bool   `json:"enable" mapstructure:"enable" yaml:"enable"`                                  //是否启用自定义日志配置
	Path             string `json:"path" mapstructure:"path" yaml:"path"`                                        //路径
	MaxSize          int    `json:"max_size" mapstructure:"max_size" yaml:"max_size"`                            //日志的最大大小（M）
	MaxBackups       int    `json:"max_backups" mapstructure:"max_backups" yaml:"max_backups"`                   //日志的最大保存数量
	MaxAge           int    `json:"max_age" mapstructure:"max_age" yaml:"max_age"`                               //日志文件存储最大天数
	Compress         bool   `json:"compress" mapstructure:"compress" yaml:"compress"`                            //是否执行压缩
	OutputMode       int    `json:"output_mode" mapstructure:"output_mode" yaml:"output_mode"`                   //输出模式 1:控制台 2：文件 3：控制台和文件都输出
	Suffix           string `json:"suffix" mapstructure:"suffix" yaml:"suffix"`                                  //日志文件后缀名
	RotationDuration int    `json:"rotation_duration" mapstructure:"rotation_duration" yaml:"rotation_duration"` //文件分隔按照时长切割
	Level            string `json:"level" mapstructure:"level" yaml:"level"`                                     // 日志等级  【trace,debug,info,notice,warn,error,fatal】
}
