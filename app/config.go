package app

import (
	"fmt"
	"github.com/spf13/viper"
	"time"
)

var Configs *configuration

type configuration struct {
	TimeOut         time.Duration `default:"1s"`
	MaxHeadingLevel int           `default:"4"`
}

func InitConfig() {
	var err error
	configFilePath := "./config.yaml"

	viper.SetConfigFile(configFilePath)
	err = viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Config file %s not found: %v", configFilePath, err))
	}

	Configs = new(configuration)
	err = viper.UnmarshalKey("configuration", Configs)
	if err != nil {
		panic(err)
	}
}
