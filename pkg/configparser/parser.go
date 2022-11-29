package configparser

import (
	"fmt"
	"github.com/spf13/viper"
)

type Parser struct {
	v *viper.Viper
}

func newViper() *viper.Viper {
	return viper.NewWithOptions()
}

func (p *Parser) UnmarshalExact(cfg interface{}) error {
	return p.v.UnmarshalExact(cfg)
}

func NewParserFromFile(fileName string) (*Parser, error) {
	v := newViper()
	v.SetConfigName(fileName)
	if err := v.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("unable to read the file %v: %w", fileName, err)
	}

	return &Parser{v: v}, nil
}
