#!/usr/bin/env bash
#可执行文件名称
SERVER_NAME="fantasy-s"
#编译二进制可执行文件
echo "build go start..."
GOOS=linux GOARCH=amd64 go build -o $SERVER_NAME
echo "build go end"
echo ""












