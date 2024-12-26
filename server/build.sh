#!/bin/bash

# 创建 build 目录
mkdir -p build

# 编译 Linux x64 版本的服务端
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build/server_linux_amd64 ./cmd 