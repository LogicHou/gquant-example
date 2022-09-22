package bootstrap

import (
	"github.com/LogicHou/gquant-example/global"

	"github.com/LogicHou/gquant/dialect"
	"go.uber.org/zap"
)

func InitializePlatform() dialect.Platform {

	platform, err := dialect.Get(&dialect.Config{
		AccessKey:  global.App.Config.Account.AccessKey,
		SecretKey:  global.App.Config.Account.SecretKey,
		Platform:   global.App.Config.Market.Platform,
		Symbol:     global.App.Config.Market.Symbol,
		Interval:   global.App.Config.Market.Interval,
		KlineRange: global.App.Config.Market.KlineRange,
	})

	if err != nil {
		global.App.Log.Fatal("cannot get dialect", zap.Error(err))
		return nil
	}

	return platform
}
