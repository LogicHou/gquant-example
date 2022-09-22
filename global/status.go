package global

import "github.com/LogicHou/gquant/indicator"

type Status struct {
	Klines         []*ReKline
	LastKline      []*ReKline
	PosAmt         float64
	PosQty         int
	EntryPrice     float64
	Leverage       float64
	PosSide        indicator.ActionType
	StopLoss       float64
	TradeStartTime string
	TradeEndTime   string
	Pid            *Pid
	CurMa5         float64
	CurMa10        float64
}

type Pid struct {
	TpPosQty    int
	CrossOffset float64
}

type ReKline struct {
	*indicator.Kline
	Ma5  float64
	Ma10 float64
}
