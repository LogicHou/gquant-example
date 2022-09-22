package middleware

import (
	"time"

	"github.com/LogicHou/gquant-example/global"

	"github.com/LogicHou/gquant"
	"github.com/LogicHou/gquant/indicator"
	"go.uber.org/zap"
)

func UpdateKline() gquant.HandlerFunc {
	return func(c *gquant.Context) {
		// 刷新Kline数据集
		if global.App.Status.Klines == nil {
			updateKline()
			c.Abort()
			return
		}
		if (c.Ticker.T - global.App.Status.LastKline[0].CloseTime) > indicator.RefreshTime[global.App.Config.Market.Interval] {
			if global.App.Status.PosAmt != 0 {
				global.App.Status.PosQty += 1
			}

			oldLastOpenTime := global.App.Status.LastKline[0].OpenTime
			updateKline()
			time.Sleep(time.Second * 3)
			for {
				if global.App.Status.LastKline[0].OpenTime == oldLastOpenTime {
					global.App.Log.Info("may klines update delay")
					time.Sleep(time.Second * 3)
					updateKline()
				} else {
					break
				}
			}
		}

		global.App.Status.CurMa5, global.App.Status.CurMa10 = curIdct(c.Ticker)

		c.Next()
	}
}

func updateKline() error {
	var err error

	klines, err := global.App.Platform.KlineRange()
	if err != nil {
		global.App.Log.Error("cannot resolve account id", zap.Error(err))
	}

	idct := indicator.New(klines)
	ma5 := idct.WithSma(5)
	ma10 := idct.WithSma(10)

	var reKlines = make([]*global.ReKline, len(klines))
	for i := 0; i < len(klines); i++ {
		reKlines[i] = &global.ReKline{klines[i], ma5[i], ma10[i]} // kline, k, d
	}

	global.App.Status.Klines = reKlines
	global.App.Status.LastKline[0] = global.App.Status.Klines[len(global.App.Status.Klines)-1]
	global.App.Status.LastKline[1] = global.App.Status.Klines[len(global.App.Status.Klines)-2]

	tmpPosAmt, tmpEntryPrice, tmpLeverage, tmpPosSide, err := global.App.Platform.PostionRisk()
	if err != nil {
		global.App.Log.Error("cannot get position risk", zap.Error(err))
	}
	if tmpPosAmt != 0 {
		global.App.Status.PosAmt, global.App.Status.EntryPrice, global.App.Status.Leverage, global.App.Status.PosSide = tmpPosAmt, tmpEntryPrice, tmpLeverage, tmpPosSide
	}

	// 有仓位的情况下重新启动，可以在这个判断条件内初始化原有的一些运行参数
	if global.App.Status.PosAmt != 0 && global.App.Status.StopLoss == 0 {
		// ...
	}

	global.App.Log.Sugar().Infof("KlineUpdated--> PosSide:%s PosAmt:%f PosQty:%d EntryPrice:%f Leverage:%f StopLoss:%f\n",
		global.App.Status.PosSide,
		global.App.Status.PosAmt,
		global.App.Status.PosQty,
		global.App.Status.EntryPrice,
		global.App.Config.Market.Leverage,
		global.App.Status.StopLoss,
	)

	return nil
}
