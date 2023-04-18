package config

type RedisOption struct {
	Network  string `yaml:"network"`  //网络类型，tcp or unix，默认tcp
	Addr     string `yaml:"addr"`     //主机名+冒号+端口，默认localhost:6379
	Password string `yaml:"password"` //密码
	DB       int64  `yaml:"db"`       // redis数据库index

	//连接池
	PoolSize     int64 `yaml:"pool_size"`      // 连接池最大socket连接数，默认为4倍CPU数， 4 * runtime.NumCPU
	MinIdleConns int64 `yaml:"min_idle_conns"` ////在启动阶段创建指定数量的Idle连接，并长期维持idle状态的连接数不少于指定数量 闲置链接数

	//超时配置
	DialTimeout  int64 `yaml:"dial_timeout"`  //连接建立超时时间，默认5秒。
	ReadTimeout  int64 `yaml:"read_timeout"`  //读超时，默认3秒， -1表示取消读超时
	WriteTimeout int64 `yaml:"write_timeout"` //写超时，默认等于读超时
	PoolTimeout  int64 `yaml:"pool_timeout"`  //当所有连接都处在繁忙状态时，客户端等待可用连接的最大等待时长，默认为读超时+1秒。

	//闲置连接检查包括IdleTimeout，MaxConnAge
	IdleCheckFrequency int64 `yaml:"idle_check_frequency"` //闲置连接检查的周期，默认为1分钟，-1表示不做周期性检查，只在客户端获取连接时对闲置连接进行处理。
	IdleTimeout        int64 `yaml:"idle_timeout"`         //闲置超时，默认5分钟，-1表示取消闲置超时检查
	MaxConnAge         int64 `yaml:"max_conn_age"`         //连接存活时长，从创建开始计时，超过指定时长则关闭连接，默认为0，即不关闭存活时长较长的连接

	//命令执行失败时的重试策略
	MaxRetries      int64 `yaml:"max_retries"`       // 命令执行失败时，最多重试多少次，默认为0即不重试
	MinRetryBackoff int64 `yaml:"min_retry_backoff"` //每次计算重试间隔时间的下限，默认8毫秒，-1表示取消间隔
	MaxRetryBackoff int64 `yaml:"max_retry_backoff"` //每次计算重试间隔时间的上限，默认512毫秒，-1表示取消间隔
}
