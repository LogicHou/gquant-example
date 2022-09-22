package bootstrap

import (
	"flag"
	"fmt"
	"os"

	"github.com/LogicHou/gquant-example/global"

	"github.com/spf13/viper"
)

var configPath = flag.String("c", "config.yaml", "config file path")

func InitializeConfig() *viper.Viper {
	flag.Parse()
	// 设置配置文件路径
	configSrc := *configPath
	// 生产环境可以通过设置环境变量来改变配置文件路径
	if configEnv := os.Getenv("VIPER_CONFIG"); configEnv != "" {
		configSrc = configEnv
	}

	// 初始化 viper
	v := viper.New()
	v.SetConfigFile(configSrc)
	v.SetConfigType("yaml")
	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("read config failed: %s \n", err))
	}

	// 将配置赋值给全局变量
	if err := v.Unmarshal(&global.App.Config); err != nil {
		fmt.Println(err)
	}

	return v
}
