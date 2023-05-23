package sarama

type KafkaOption struct {
	TLSOption  KafkaTLSOption  `json:"tls_option" mapstructure:"tls_option"  yaml:"tls_option"`
	CAOption   KafkaCAOption   `json:"ca_option" mapstructure:"ca_option"  yaml:"ca_option"`
	SASLOption KafkaSASLOption `json:"sasl_option" mapstructure:"sasl_option"  yaml:"sasl_option"`
	Timeout    int             `json:"timeout" mapstructure:"timeout"  yaml:"timeout"`
	Broker     []string        `json:"broker" mapstructure:"broker"  yaml:"broker"`
}

type KafkaTLSOption struct {
	Enable   bool   `json:"enable" mapstructure:"enable" yaml:"enable"`
	CertPath string `json:"cert_path" mapstructure:"cert_path" yaml:"cert_path"`
	KeyPath  string `json:"key_path" mapstructure:"key_path" yaml:"key_path"`
}

type KafkaCAOption struct {
	Enable bool   `json:"enable" mapstructure:"enable" yaml:"enable"`
	CAPath string `json:"ca_path" mapstructure:"ca_path" yaml:"ca_path"`
}

type KafkaSASLOption struct {
	Enable   bool   `json:"enable" mapstructure:"enable" yaml:"enable"`
	User     string `json:"user" mapstructure:"user" yaml:"user"`
	Password string `json:"password" mapstructure:"password" yaml:"password"`
}
