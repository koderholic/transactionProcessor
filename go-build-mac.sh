#!/bin/bash
#go build  -o stib_mac.app -ldflags "-s -w" && mv stib_mac.app app/.
#go build  -o stib_mac.app -ldflags "-s -w" && upx "-9" stib_mac.app && mv stib_mac.app app/.
go build  -o stib_mac.app -ldflags "-s -w"
