package olivere

type ClientOption struct {
	URL         []string `json:"url" mapstructure:"url" yaml:"url"`
	Index       string   `json:"index" mapstructure:"index" yaml:"index"`
	Username    string   `json:"username" mapstructure:"username" yaml:"username"`
	Password    string   `json:"password" mapstructure:"password" yaml:"password"`
	Shards      int      `json:"shards" mapstructure:"shards" yaml:"shards"`
	Replicas    int      `json:"replicas" mapstructure:"replicas" yaml:"replicas"`
	Sniff       *bool    `json:"sniff" mapstructure:"sniff" yaml:"sniff"`
	Healthcheck *bool    `json:"healthcheck" mapstructure:"healthcheck" yaml:"healthcheck"`
	InfoLog     string   `json:"info_log" mapstructure:"info_log" yaml:"info_log"`
	ErrorLog    string   `json:"error_log" mapstructure:"error_log" yaml:"error_log"`
	TraceLog    string   `json:"trace_log" mapstructure:"trace_log" yaml:"trace_log"`
}
