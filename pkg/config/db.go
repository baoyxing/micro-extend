package config

type DBOption struct {
	Driver         string `json:"driver" mapstructure:"driver" yaml:"driver"`
	Source         string `json:"source" mapstructure:"source" yaml:"source"`
	SqlLog         bool   `json:"sql_log" mapstructure:"sql_log" yaml:"sql_log"`
	MaxOpenCons    int    `json:"max_open_cons" mapstructure:"max_open_cons" yaml:"max_open_cons"`
	MaxIdleCons    int    `json:"max_idle_cons" mapstructure:"max_idle_cons" yaml:"max_idle_cons"`
	ConMaxLifetime int    `json:"con_max_lifetime" mapstructure:"con_max_lifetime" yaml:"con_max_lifetime"`
}
