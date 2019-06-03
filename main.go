package main

import (
	"flag"
	"fmt"
)

func main() {

	var loadConfig string

	flag.StringVar(&loadConfig, "config", ".", "directory location of config file")
	flag.Parse()

	if loadConfig == "" {
		fmt.Printf("----------------------------------------------------------\n")
		flag.Usage()
		fmt.Printf("----------------------------------------------------------\n")
	}

	initializeConf(loadConfig)
	startProcess()
}
