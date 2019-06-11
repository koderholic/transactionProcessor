package main

import (
	"fmt"
	"os"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type conf struct {
	InDir       string `yaml:"inDir"`
	OutDir      string `yaml:"outDir"`
	BackupDir   string `yaml:"backupDir"`
	LogDir      string `yaml:"logDir"`
	Suffix      string `yaml:"suffix"`
	Keyword     string `yaml:"keyword"`
	TrimerIndex string `yaml:"trimerIndex"`
}

var config conf

func logToFile(log string) {
	// If the file doesn't exist, create it, or append to the file
	fileName := fmt.Sprintf("%s%slogs_%v.txt", config.LogDir, string(os.PathSeparator), time.Now().Format("20060102"))
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		println("ERROR", err.Error())
	}
	if _, err := file.Write([]byte(log)); err != nil {
		println("ERROR", err.Error())
	}
	if err := file.Close(); err != nil {
		println("ERROR", err.Error())
	}
}

func (c *conf) init(configDir string) {

	dir, dirErr := os.Getwd()
	if dirErr != nil {
		fmt.Printf("Cannot set default input/output directory to the current working directory >> %s", dirErr)
		logToFile(fmt.Sprintf("\n\nCannot set default input/output directory to the current working directory >> %s \n", dirErr))
	}
	print(dir)

	viper.SetConfigName("conf")
	viper.AddConfigPath(".")
	viper.AddConfigPath(dir)
	viper.AddConfigPath(configDir)
	viper.SetDefault("inDir", dir)
	viper.SetDefault("outDir", dir)
	viper.SetDefault("backupDir", "")
	viper.SetDefault("logDir", dir)
	viper.SetDefault("suffix", ".summary")
	viper.SetDefault("keyword", "48=")
	viper.SetDefault("trimerIndex", "8=FIX.4.4")
	viper.WatchConfig()

	err := viper.ReadInConfig()
	if err != nil {
		logToFile(fmt.Sprintf("\n\nfatal error: could not read from config file >>%s \n", err))
		panic(fmt.Errorf("\n fatal error: could not read from config file >>%s ", err))
	}

	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
		err := viper.ReadInConfig()
		if err != nil {
			fmt.Printf("\n fatal error: could not read from config file >>%s ", err)
			logToFile(fmt.Sprintf("\n\n\n fatal error: could not read from config file >>%s\n ", err))
		}
		viper.Unmarshal(c)
	})

	viper.Unmarshal(c)
}

func initializeConf(configDir string) {
	config.init(configDir)
}
