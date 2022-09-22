package middleware

import (
	"time"

	"github.com/LogicHou/gquant-example/global"

	"github.com/LogicHou/gquant"
	"github.com/LogicHou/gquant/indicator"
	"github.com/LogicHou/gquant/utils"
	"go.uber.org/zap"
)

func StopLoss() gquant.HandlerFunc {
	return func(c *gquant.Context) {
		// 止损逻辑
		if stCondition(c.Ticker) {
			global.App.Status.TradeEndTime = utils.MsToTime(time.Now().UnixMilli())
			global.App.Log.Sugar().Infof("ST - Action: %s InTime: %s OutTime: %s InPrice: %f OutPrice: %f Ratio: %f", global.App.Status.PosSide, global.App.Status.TradeStartTime, global.App.Status.TradeEndTime, global.App.Status.EntryPrice, c.Ticker.C, calcProfit(c.Ticker))

			err := closePosition(c.Ticker)
			if err != nil {
				global.App.Log.Error("cannot close position in stop loss", zap.Error(err))
				c.Abort()
				return
			}
		}
		c.Next()
	}
}

func stCondition(t *indicator.Ticker) bool {
	switch global.App.Status.PosSide {
	case indicator.ActionBuy:
		if t.C < global.App.Status.StopLoss {
			global.App.Log.Sugar().Infof("BUY - ST condition triggered - PosSide:%s ticker.C:%f StopLoss:%f\n", global.App.Status.PosSide, t.C, global.App.Status.StopLoss)
			return true
		}
	case indicator.ActionSell:
		if t.C > global.App.Status.StopLoss {
			global.App.Log.Sugar().Infof("SELL - ST condition triggered - PosSide:%s ticker.C:%f StopLoss:%f\n", global.App.Status.PosSide, t.C, global.App.Status.StopLoss)
			return true
		}
	}

	return false
}
