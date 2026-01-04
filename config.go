package seedgo

import (
	"github.com/spf13/viper"
)

func Init() {
	loadConfig()
	ParseSystemConfig()
	InitLogger()
	if ServerConfig.Debug {
		Logger.Info(viper.AllSettings())
	}

	ParseDatabaseConfig()
	ParseRedisConf()
}

func loadConfig() {
	viper.SetConfigFile(*configFile)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

}
