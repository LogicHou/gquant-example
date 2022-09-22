package main

import (
	"github.com/LogicHou/gquant-example/bootstrap"
	"github.com/LogicHou/gquant-example/global"
	"github.com/LogicHou/gquant-example/middleware"

	"github.com/LogicHou/gquant"
)

func main() {
	bootstrap.InitializeConfig()
	global.App.Log = bootstrap.InitializeLog()
	global.App.Platform = bootstrap.InitializePlatform()
	global.App.Status = bootstrap.InitializeStatus()

	r := gquant.New()
	r.Use(gquant.Logger())
	r.Use(gquant.Recovery())
	r.Use(middleware.UpdateKline())
	r.Use(middleware.OpenPosition())
	r.Use(middleware.TakeProfit())
	r.Use(middleware.StopLoss())

	r.AddHandle(func(c *gquant.Context) {
		// println(c.Ticker.C, c.Ticker.T)
	})

	r.Run(global.App.Platform)
}
