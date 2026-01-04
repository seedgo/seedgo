package seedgo

import "github.com/spf13/viper"

type ServerConfigType struct {
	Port  int
	Debug bool
}

var ServerConfig ServerConfigType

func ParseSystemConfig() {
	viper.SetDefault("server.port", 10016)
	viper.SetDefault("server.debug", true)

	ServerConfig.Port = viper.GetInt("server.port")
	ServerConfig.Debug = viper.GetBool("server.debug")
}
