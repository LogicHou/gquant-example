package global

import (
	"github.com/LogicHou/gquant-example/config"

	"github.com/LogicHou/gquant/dialect"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type Application struct {
	ConfigViper *viper.Viper
	Config      config.Configuration
	Log         *zap.Logger
	Platform    dialect.Platform
	Status      *Status
}

var App = new(Application)
