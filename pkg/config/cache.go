package config

type RedisOption struct {
	Network  string `json:"network" mapstructure:"network" yaml:"network"`    //网络类型，tcp or unix，默认tcp
	Addr     string `json:"addr" mapstructure:"addr" yaml:"addr"`             //主机名+冒号+端口，默认localhost:6379
	Username string `json:"username" mapstructure:"username" yaml:"username"` // 用户名
	Password string `json:"password" mapstructure:"password" yaml:"password"` //密码
	DB       int64  `json:"db" mapstructure:"db" yaml:"db"`                   // redis数据库index

	//连接池
	PoolSize    int64 `json:"pool_size" mapstructure:"pool_size" yaml:"pool_size"`             // 连接池最大socket连接数，默认为4倍CPU数， 4 * runtime.NumCPU
	MinIdleCons int64 `json:"min_idle_cons" mapstructure:"min_idle_cons" yaml:"min_idle_cons"` ////在启动阶段创建指定数量的Idle连接，并长期维持idle状态的连接数不少于指定数量 闲置链接数

	//超时配置
	DialTimeout  int64 `json:"dial_timeout" mapstructure:"dial_timeout" yaml:"dial_timeout"`    //连接建立超时时间，默认5秒。
	ReadTimeout  int64 `json:"read_timeout" mapstructure:"read_timeout" yaml:"read_timeout"`    //读超时，默认3秒， -1表示取消读超时
	WriteTimeout int64 `json:"write_timeout" mapstructure:"write_timeout" yaml:"write_timeout"` //写超时，默认等于读超时
	PoolTimeout  int64 `json:"pool_timeout" mapstructure:"pool_timeout" yaml:"pool_timeout"`    //当所有连接都处在繁忙状态时，客户端等待可用连接的最大等待时长，默认为读超时+1秒。

	//闲置连接检查包括IdleTimeout，MaxConnAge
	IdleCheckFrequency int64 `json:"idle_check_frequency" mapstructure:"idle_check_frequency" yaml:"idle_check_frequency"` //闲置连接检查的周期，默认为1分钟，-1表示不做周期性检查，只在客户端获取连接时对闲置连接进行处理。
	IdleTimeout        int64 `json:"idle_timeout" mapstructure:"idle_timeout" yaml:"idle_timeout"`                         //闲置超时，默认5分钟，-1表示取消闲置超时检查
	MaxConnAge         int64 `json:"max_conn_age" mapstructure:"max_conn_age" yaml:"max_conn_age"`                         //连接存活时长，从创建开始计时，超过指定时长则关闭连接，默认为0，即不关闭存活时长较长的连接

	//命令执行失败时的重试策略
	MaxRetries      int64 `json:"max_retries" mapstructure:"max_retries" yaml:"max_retries"`                   // 命令执行失败时，最多重试多少次，默认为0即不重试
	MinRetryBackoff int64 `json:"min_retry_backoff" mapstructure:"min_retry_backoff" yaml:"min_retry_backoff"` //每次计算重试间隔时间的下限，默认8毫秒，-1表示取消间隔
	MaxRetryBackoff int64 `json:"max_retry_backoff" mapstructure:"max_retry_backoff" yaml:"max_retry_backoff"` //每次计算重试间隔时间的上限，默认512毫秒，-1表示取消间隔
}
