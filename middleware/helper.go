package middleware

import (
	"fmt"

	"github.com/LogicHou/gquant-example/global"

	"github.com/LogicHou/gquant/indicator"
	"github.com/LogicHou/gquant/utils"
)

// 计算当前指标
func curIdct(t *indicator.Ticker) (float64, float64) {
	sklen := len(global.App.Status.Klines)
	klines := make([]*indicator.Kline, sklen+1)
	for i := 0; i < sklen; i++ {
		klines[i] = global.App.Status.Klines[i].Kline
	}
	klines[sklen] = &indicator.Kline{
		OpenTime:  t.S,
		CloseTime: t.E,
		Open:      t.O,
		High:      t.H,
		Low:       t.L,
		Close:     t.C,
		Volume:    t.V,
	}
	idct := indicator.New(klines)
	curMa5 := idct.WithSma(5)
	curMa10 := idct.WithSma(10)
	return curMa5[len(curMa5)-1], curMa10[len(curMa10)-1]
}

// 计算开仓数量
func calcMqrginQty(t *indicator.Ticker) (float64, error) {
	if global.App.Config.Market.MarginRatio > 0 {
		res, err := global.App.Platform.GetAccountInfo()
		if err != nil {
			return 0.0, err
		}
		if utils.StrToF64(res.TotalWalletBalance) < global.App.Config.Market.MarginFloor {
			return 0.0, fmt.Errorf("TotalWalletBalance less than MarginFloor: %v", err)
		}
		return utils.FRound((global.App.Config.Market.MarginRatio / 100.00 * utils.StrToF64(res.TotalWalletBalance)) * global.App.Config.Market.Leverage / t.C), nil
	}
	if global.App.Config.Market.Margin > 0 {
		return utils.FRound(global.App.Config.Market.Margin * global.App.Config.Market.Leverage / t.C), nil
	}

	return 0.0, nil
}

// 计算前低
func findFrontLow(t *indicator.Ticker) (float64, error) {
	if global.App.Status.PosSide == indicator.ActionBuy {
		low := t.L
		for i := len(global.App.Status.Klines) - 1; i > 0; i-- {
			if global.App.Status.Klines[i].Low < low && global.App.Status.Klines[i].Low < global.App.Status.Klines[i-1].Low {
				return global.App.Status.Klines[i].Low, nil
			}
		}
	}

	if global.App.Status.PosSide == indicator.ActionSell {
		high := t.H
		for i := len(global.App.Status.Klines) - 1; i > 0; i-- {
			if global.App.Status.Klines[i].High > high && global.App.Status.Klines[i].High > global.App.Status.Klines[i-1].High {
				return global.App.Status.Klines[i].High, nil
			}
		}
	}

	return 0.0, fmt.Errorf("not found stoploss condition")
}

func calcProfit(t *indicator.Ticker) float64 {
	const feeRatio float64 = 0.08
	if global.App.Status.PosSide == indicator.ActionBuy {
		return (t.C/global.App.Status.EntryPrice-1)*100 - feeRatio
	}
	return (t.C/global.App.Status.EntryPrice-1)*-100 - feeRatio
}

// 平仓函数
func closePosition(t *indicator.Ticker) error {
	global.App.Log.Info("close position!")
	// 调用平台API进行平仓
	// err := global.App.Platform.CloseMarketPosition(global.App.Status.PosAmt)
	// if err != nil {
	// 	return err
	// }

	global.App.Status.EntryPrice = 0
	global.App.Status.PosAmt = 0
	global.App.Status.StopLoss = 0
	global.App.Status.PosQty = 0
	global.App.Status.TradeStartTime = ""
	global.App.Status.TradeEndTime = ""

	return nil
}
