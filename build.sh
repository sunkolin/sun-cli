#!/bin/bash

# windows
GOOS=windows GOARCH=amd64 go build -o sun.exe main.go

# mac intel
GOOS=darwin GOARCH=amd64 go build -o sun main.go

# mac arm
GOOS=darwin GOARCH=arm64 go build -o sun main.go
