package configparser

import "github.com/spf13/pflag"

const (
	configFlagName = "config"
)

var configFlag *string

func Flags(flags *pflag.FlagSet) {
	configFlag = flags.String(configFlagName, "", "Path to the config file")
}

func getConfigFlag() string {
	return *configFlag
}
