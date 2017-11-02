#!/bin/bash
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o dist-linux/datawasherLinux app.go
