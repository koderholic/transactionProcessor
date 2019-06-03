package main

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type conf struct {
	Suffix  string `yaml:"suffix"`
	Keyword string `yaml:"keyword"`
}

var config conf

func (c *conf) getConf() {

	viper.SetConfigName("conf")
	viper.AddConfigPath(".")
	viper.SetDefault("suffix", ".summary")
	viper.SetDefault("keyword", "48=")
	viper.WatchConfig()

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("\n fatal error: could not read from config file >>%s ", err))
	}

	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
		err := viper.ReadInConfig()
		if err != nil {
			panic(fmt.Errorf("\n fatal error: could not read from config file >>%s ", err))
		}
		viper.Unmarshal(c)
	})

	viper.Unmarshal(c)
}

func initializeConf() {
	go config.getConf()
}
