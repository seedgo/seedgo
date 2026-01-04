package seedgo

import "flag"

var configFile *string

func InitCmd() {
	configFile = flag.String("config", "./config/application.yaml", "config file path")
	flag.Parse()
}
