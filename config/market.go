package config

type Market struct {
	Platform    string  `mapstructure:"platform" json:"platform" yaml:"platform"`
	Symbol      string  `mapstructure:"symbol" json:"symbol" yaml:"symbol"`
	Interval    string  `mapstructure:"interval" json:"interval" yaml:"interval"`
	Leverage    float64 `mapstructure:"leverage" json:"leverage" yaml:"leverage"`
	KlineRange  int     `mapstructure:"kline_range" json:"kline_range" yaml:"kline_range"`
	Margin      float64 `mapstructure:"margin" json:"margin" yaml:"margin"`
	MarginRatio float64 `mapstructure:"margin_ratio" json:"margin_ratio" yaml:"margin_ratio"`
	MarginFloor float64 `mapstructure:"margin_floor" json:"margin_floor" yaml:"margin_floor"`
}
