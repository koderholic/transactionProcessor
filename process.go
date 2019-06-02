package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"github.com/fsnotify/fsnotify"
)

/*
1 Listen for (create and write)Changes on a specific file(*.summary) in the input directory
2 Once they is a change, Read current content of the File
3 Create a slice with the Content
4 Identify specific transaction data in the content slice to be removed
	4.1 Create new Unique file
	4.2 Move each identified transaction data into the unique file
5 Remove identified transaction data from content slice, creating a new slice
6 convert new slice back to string
7 save new slice back to the watched file in the input directory

TODO :
1. Build for multiple platform
2. Run task scheduler on windows
3. Add config file for double click to start with viper package
4. Write test
*/

func startProcess(inDir, outDir, backupDir string) {
	if _, err := os.Stat(inDir); os.IsNotExist(err) {
		fmt.Println("Input directory does not exists")
		return
	}

	if _, err := os.Stat(outDir); os.IsNotExist(err) {
		fmt.Println("Output directory does not exists")
		return
	}

	if backupDir != "" {
		if _, err := os.Stat(backupDir); os.IsNotExist(err) {
			fmt.Println("Backup directory does not exists")
			return
		}
	}

	//start monitoring input directory
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		fmt.Println("ERROR", err)
	}
	defer watcher.Close()

	tick := time.Tick(time.Second)
	eventMap := make(map[string]fsnotify.Event)

	go func() {
		type fileNamer struct {
			day   time.Weekday
			count int
		}
		for {
			fileNameCounter := fileNamer{time.Now().Weekday(), 0}
			select {
			case <-tick:
				for name, event := range eventMap {
					delete(eventMap, name)

					var newFileSlice []string
					fileSlice := fileToSlice(event.Name)

					for _, content := range fileSlice {
						if !strings.Contains(content, "transactionToLookOutFor") {
							newFileSlice = append(newFileSlice, content)
							continue
						}
						println("Found Transaction:", content[:64])

						transactionLine := strings.Split(content, "\n")
						fileName := fmt.Sprintf("STLB_Transact_%v_%v_%v.txt", time.Now().Format("20060102"), time.Now().Format("150405"), fileNameCounter.count)

						sliceToFile(transactionLine, outDir+string(os.PathSeparator)+fileName)
						println("Created Unique Transaction File: ", fileName)

						if backupDir != "" {
							backupToFile(content, fmt.Sprintf("%s%sbackup_%v.txt", backupDir, string(os.PathSeparator), time.Now().Format("20060102")))
						}

						if today := &fileNameCounter; today.day == time.Now().Weekday() {
							today.count++
						} else {
							today.day = time.Now().Weekday()
							today.count = 0
						}
						//Proces New Transaction to a New File
					}

					sliceToFile(newFileSlice, event.Name)

				}

			// watch for events
			case event := <-watcher.Events:
				if strings.Contains(event.Name, "summary") {
					switch {
					case event.Op&fsnotify.Write == fsnotify.Write,
						event.Op&fsnotify.Create == fsnotify.Create:
						eventMap[event.Name] = event
					}
				}

			// watch for errors
			case err := <-watcher.Errors:
				fmt.Println("ERROR:", err)

			}
		}
	}()

	if err := watcher.Add(inDir); err != nil {
		fmt.Println("FATAL ERROR:", err)
		os.Exit(0)
	}

	//This section exits the process
	var exit string
	fmt.Print("Type 'exit' and press 'Enter' to stop process...")
	for {
		if fmt.Scanln(&exit); exit == "exit" {
			os.Exit(0)
		}
	}
}

func fileToSlice(fileName string) (fileSlice []string) {
	fileBytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		return
	}

	fileString := string(fileBytes)

	fileSlice = strings.Split(fileString, "\n")
	return
}

func sliceToFile(fileSlice []string, fileName string) {
	fileString := strings.Join(fileSlice, "\n")
	ioutil.WriteFile(fileName, []byte(fileString), 666)
}

func backupToFile(transactionLine string, fileName string) {
	// If the file doesn't exist, create it, or append to the file
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		println("ERROR", err.Error())
	}
	if _, err := file.Write([]byte(transactionLine)); err != nil {
		println("ERROR", err.Error())
	}
	if err := file.Close(); err != nil {
		println("ERROR", err.Error())
	}
}
