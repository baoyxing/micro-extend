package hdfs

type ClientOption struct {
	Addresses []string `yaml:"addresses"`
	User      string   `yaml:"user"`
}
