package config

type Configuration struct {
	Account Account `mapstructure:"account" json:"account" yaml:"account"`
	Market  Market  `mapstructure:"market" json:"market" yaml:"market"`
	Pid     Pid     `mapstructure:"pid" json:"pid" yaml:"pid"`
}
