package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	var in, out, b string

	flag.StringVar(&out, "out", ".", "Output Directory of transactions '.unique' files")
	flag.StringVar(&in, "in", ".", "Input Directory of transactions '.summary' files")
	flag.StringVar(&b, "b", "", "Backup Directory of transactions '.summary' files")
	flag.Parse()

	if in == "" && out == "" {
		fmt.Printf("----------------------------------------------------------\n")
		flag.Usage()
		fmt.Printf("----------------------------------------------------------\n")

		os.Exit(0)
	}

	initializeConf()
	startProcess(in, out, b)
}
