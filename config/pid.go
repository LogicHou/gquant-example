package config

type Pid struct {
	TpPosQty    int     `mapstructure:"tp_pos_qty" json:"tp_pos_qty" yaml:"tp_pos_qty"`
	CrossOffset float64 `mapstructure:"cross_offset" json:"cross_offset" yaml:"cross_offset"`
}
