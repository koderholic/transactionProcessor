#!/bin/bash
#GOOS=windows GOARCH=386 go build -o build/stib_app.exe -ldflags "-s -w" && upx stib_win.exe && mv stib_win.exe app/.
env GOOS=windows GOARCH=amd64 go build -o build/stib_app.exe
#go build -o build/stib_app.exe