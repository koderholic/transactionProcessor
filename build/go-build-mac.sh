#!/bin/bash
#go build  -o stib_mac.app -ldflags "-s -w" && mv stib_mac.app app/.
go build  -o build/stib_mac.app -ldflags "-s -w"

function pause(){
   read -p "$*"
}
pause 'Press [Enter] key to continue...'
