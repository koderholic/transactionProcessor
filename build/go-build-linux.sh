#!/bin/bash
env GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o build/t24Fix_interface.elf -ldflags "-s -w" && upx stib_linux.elf

#function pause(){
#   read -p "$*"
#}
#pause 'Press [Enter] key to continue...'