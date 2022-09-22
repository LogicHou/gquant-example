package middleware

import (
	"time"

	"github.com/LogicHou/gquant-example/global"

	"github.com/LogicHou/gquant"
	"github.com/LogicHou/gquant/indicator"
	"github.com/LogicHou/gquant/utils"
	"go.uber.org/zap"
)

func TakeProfit() gquant.HandlerFunc {
	return func(c *gquant.Context) {
		// 止盈逻辑
		if global.App.Status.PosQty > global.App.Config.Pid.TpPosQty {
			if tpCondition(c.Ticker) {
				global.App.Status.TradeEndTime = utils.MsToTime(time.Now().UnixMilli())
				global.App.Log.Sugar().Infof("TP - Action: %s InTime: %s OutTime: %s InPrice: %f OutPrice: %f Ratio: %f", global.App.Status.PosSide, global.App.Status.TradeStartTime, global.App.Status.TradeEndTime, global.App.Status.EntryPrice, c.Ticker.C, calcProfit(c.Ticker))

				err := closePosition(c.Ticker)
				if err != nil {
					global.App.Log.Error("cannot close position in take profit", zap.Error(err))
					c.Abort()
					return
				}
			}
			c.Abort()
			return
		}
		c.Next()
	}
}

func tpCondition(t *indicator.Ticker) bool {
	switch global.App.Status.PosSide {
	case indicator.ActionBuy:
		if (t.C/global.App.Status.EntryPrice - 1) > 0.05 {
			global.App.Log.Info("BUY - TP > 5% TP condition triggered")
			return true
		}
	case indicator.ActionSell:
		if (t.C/global.App.Status.EntryPrice-1)*-1 > 0.05 {
			global.App.Log.Info("SELL - TP > 5% TP condition triggered")
			return true
		}
	}

	return false
}
