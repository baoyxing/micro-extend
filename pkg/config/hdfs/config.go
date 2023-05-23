package hdfs

type ClientOption struct {
	Addresses []string `json:"addresses" mapstructure:"addresses" yaml:"addresses"`
	User      string   `json:"user" mapstructure:"user" yaml:"user"`
}
