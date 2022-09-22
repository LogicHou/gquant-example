package middleware

import (
	"time"

	"github.com/LogicHou/gquant-example/global"

	"github.com/LogicHou/gquant"
	"github.com/LogicHou/gquant/indicator"
	"github.com/LogicHou/gquant/utils"
	"go.uber.org/zap"
)

func OpenPosition() gquant.HandlerFunc {
	return func(c *gquant.Context) {
		if global.App.Status.PosAmt == 0 {
			if global.App.Status.CurMa5 < global.App.Status.CurMa10 {
				global.App.Status.PosSide = indicator.ActionBuy
			} else if global.App.Status.CurMa5 > global.App.Status.CurMa10 {
				global.App.Status.PosSide = indicator.ActionSell
			}

			if openCondition() {
				qty, err := calcMqrginQty(c.Ticker)
				if err != nil {
					global.App.Log.Error("cannot get margin qty", zap.Error(err))
					c.Abort()
					return
				}
				frontLow, err := findFrontLow(c.Ticker)
				if err != nil {
					global.App.Log.Error("cannot get findFrontLow", zap.Error(err))
					c.Abort()
					return
				}
				global.App.Status.StopLoss = frontLow

				global.App.Log.Info("create order!")
				// 调用平台API进行开仓
				// err = global.App.Platform.CreateMarketOrder(global.App.Status.PosSide, qty, global.App.Status.StopLoss)
				// if err != nil {
				// 	global.App.Log.Error("cannot create market order", zap.Error(err))
				// 	c.Abort()
				// 	return
				// }

				global.App.Status.PosAmt = qty
				global.App.Status.EntryPrice = c.Ticker.C
				global.App.Status.TradeStartTime = utils.MsToTime(time.Now().UnixMilli())

				if err != nil {
					global.App.Log.Error("cannot calcMovingStop", zap.Error(err))
				}

				global.App.Log.Sugar().Infof("OP - Action: %s  EntryPrice: %f STOPLOSS: %f PosAmt: %f", global.App.Status.PosSide, global.App.Status.EntryPrice, global.App.Status.StopLoss, global.App.Status.PosAmt)
				time.Sleep(time.Second * 3)
			}
			c.Abort()
			return
		}

		c.Next()
	}
}

func openCondition() bool {
	switch global.App.Status.PosSide {
	case indicator.ActionBuy:
		if global.App.Status.LastKline[0].Ma5 < global.App.Status.LastKline[0].Ma10 && global.App.Status.CurMa5 > global.App.Status.CurMa10+global.App.Config.Pid.CrossOffset {
			global.App.Log.Info("BUY - ma5 up cross ma10 open condition triggered")
			return true
		}
	case indicator.ActionSell:
		if global.App.Status.LastKline[0].Ma5 > global.App.Status.LastKline[0].Ma10 && global.App.Status.CurMa5 < global.App.Status.CurMa10-global.App.Config.Pid.CrossOffset {
			global.App.Log.Info("SELL - ma5 down cross ma10 open condition triggered")
			return true
		}
	}

	return false
}
