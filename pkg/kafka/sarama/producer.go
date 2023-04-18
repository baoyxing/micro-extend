package sarama

import (
	"crypto/tls"
	"crypto/x509"
	shSarama "github.com/Shopify/sarama"
	saramaConf "github.com/baoyxing/micro-extend/pkg/config/kafka/sarama"
	"io/ioutil"
	"sync"
	"time"
)

type KafkaProducer struct {
	shSarama.SyncProducer
	sync.RWMutex
}

func createTlsConfig(opt *saramaConf.KafkaOption) (*tls.Config, error) {
	t := &tls.Config{InsecureSkipVerify: true}
	if opt.TLSOption.Enable {
		cert, err := tls.LoadX509KeyPair(opt.TLSOption.CertPath, opt.TLSOption.KeyPath)
		if err != nil {
			return nil, err
		}
		t.Certificates = []tls.Certificate{cert}
	}

	if opt.CAOption.Enable {
		caCert, err := ioutil.ReadFile(opt.CAOption.CAPath)
		if err != nil {
			return nil, err
		}
		caCertPool := x509.NewCertPool()
		caCertPool.AppendCertsFromPEM(caCert)
		t.RootCAs = caCertPool
	}
	return t, nil
}
func NewKafkaProducer(opt *saramaConf.KafkaOption) (*KafkaProducer, error) {
	tlsConfig, err := createTlsConfig(opt)
	if err != nil {
		return nil, err
	}
	config := shSarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.Timeout = time.Duration(opt.Timeout) * time.Second
	config.Producer.RequiredAcks = shSarama.WaitForAll
	if opt.TLSOption.Enable || opt.CAOption.Enable {
		config.Net.TLS.Enable = true
	}
	config.Net.TLS.Config = tlsConfig
	if opt.SASLOption.Enable {
		config.Net.SASL.Enable = true
		config.Net.SASL.Handshake = true
		config.Net.SASL.User = opt.SASLOption.User
		config.Net.SASL.Password = opt.SASLOption.Password
	}
	producer, err := shSarama.NewSyncProducer(opt.Broker, config)
	if err != nil {
		return nil, err
	}
	return &KafkaProducer{
		SyncProducer: producer,
		RWMutex:      sync.RWMutex{},
	}, nil
}

func (p *KafkaProducer) Close() error {
	p.RWMutex.Lock()
	defer p.RWMutex.RLock()
	return p.SyncProducer.Close()
}

func (p *KafkaProducer) Send(topic, content string) error {
	p.RWMutex.RLock()
	p.RWMutex.RUnlock()
	_, _, err := p.SendMessage(&shSarama.ProducerMessage{
		Topic: topic,
		Value: shSarama.StringEncoder(content),
	})
	return err
}

func (p *KafkaProducer) SendN(topic string, contents []string) error {
	p.RWMutex.RLock()
	p.RWMutex.RUnlock()
	var msgs []*shSarama.ProducerMessage
	for _, value := range contents {
		msgs = append(msgs, &shSarama.ProducerMessage{
			Topic: topic,
			Value: shSarama.StringEncoder(value),
		})
	}
	return p.SendMessages(msgs)
}
