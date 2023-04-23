package olivere

type ClientOption struct {
	URL         string `yaml:"url"`
	Index       string `yaml:"index"`
	Username    string `yaml:"username"`
	Password    string `yaml:"password"`
	Shards      int    `yaml:"shards"`
	Replicas    int    `yaml:"replicas"`
	Sniff       *bool  `yaml:"sniff"`
	Healthcheck *bool  `yaml:"healthcheck"`
	Infolog     string `yaml:"infolog"`
	Errorlog    string `yaml:"errorlog"`
	Tracelog    string `yaml:"tracelog"`
}
