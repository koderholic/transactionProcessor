#!/bin/bash
env GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o build/app.elf -ldflags "-s -w" && upx app.elf