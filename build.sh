#!/bin/bash

# 创建 build 目录
mkdir -p build

# 编译不同平台的版本
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o build/client_windows_amd64.exe ./client/cmd 