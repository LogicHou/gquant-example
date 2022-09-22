package bootstrap

import (
	"github.com/LogicHou/gquant-example/global"
)

func InitializeStatus() *global.Status {

	return &global.Status{
		LastKline:  make([]*global.ReKline, 2),
		PosAmt:     0.00,
		PosQty:     0,
		EntryPrice: 0.00,
		PosSide:    "",
		StopLoss:   0.00,
		CurMa5:     0.00,
		CurMa10:    0.00,
		Pid: &global.Pid{
			TpPosQty:    global.App.Config.Pid.TpPosQty,
			CrossOffset: global.App.Config.Pid.CrossOffset,
		},
	}

}
