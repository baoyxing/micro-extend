package sarama

type KafkaOption struct {
	TLSOption  KafkaTLSOption  `yaml:"tls_option"`
	CAOption   KafkaCAOption   `yaml:"ca_option"`
	SASLOption KafkaSASLOption `yaml:"sasl_option"`
	Timeout    int             `yaml:"timeout"`
	Broker     []string        `yaml:"broker"`
}

type KafkaTLSOption struct {
	Enable   bool   `yaml:"enable"`
	CertPath string `yaml:"cert_path"`
	KeyPath  string `yaml:"key_path"`
}

type KafkaCAOption struct {
	Enable bool   `yaml:"enable"`
	CAPath string `yaml:"ca_path"`
}

type KafkaSASLOption struct {
	Enable   bool   `yaml:"enable"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}
