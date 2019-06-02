#!/bin/bash
GOOS=windows GOARCH=386 go build  -o stib_win.exe -ldflags "-s -w" && upx stib_win.exe && mv stib_win.exe app/.
