#!/bin/bash

# 创建 build 目录
mkdir -p build

# 编译 Linux x64 版本的服务端
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build/client_linux_amd64 ./cmd 
CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -o build/client_darwin_arm64 ./cmd 
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o build/client_darwin_amd64 ./cmd 
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o build/client_windows_amd64.exe ./cmd