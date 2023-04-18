package config

type DBOption struct {
	Driver          string `yaml:"driver"`
	Source          string `yaml:"source"`
	SqlLog          bool   `yaml:"sql_log"`
	MaxOpenConns    int    `yaml:"max_open_conns"`
	MaxIdleConns    int    `yaml:"max_idle_conns"`
	ConnMaxLifetime int    `yaml:"conn_max_lifetime"`
}
