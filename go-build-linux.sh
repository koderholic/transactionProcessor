#!/bin/bash
#upx stib_linux.elf &&
CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o stib_linux.elf -ldflags "-s -w" && upx stib_linux.elf
mv 

# $env:GOOS = "linux" $env:GOARCH = "amd64" go build -o stibelf.exe
