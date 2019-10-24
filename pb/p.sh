#!/usr/bin/env bash

#过滤rpc_*.proto
file=`find .  -type  f |grep '.proto' |grep -v rpc_*.proto`
echo $file

protoc  --go_out=. $file



