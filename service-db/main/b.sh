#!/usr/bin/env bash

#mysql 连接地址
MYSQL_ADDR="root:123456@tcp(192.168.0.112:3306)/account_db?charset=utf8"

#xorm 模板路径
#XORM_MODULE_PATH="/f/unity/fantacy/server/src/github.com/Irfish/component/xorm/templates"
XORM_MODULE_PATH="/D/git/fantacy/server\src/github.com/Irfish/component/xorm/templates"


#导出go文件存放路径
cd ..
OUT_PATH=`pwd`

#将table 映射到 go
xorm reverse mysql "$MYSQL_ADDR" $XORM_MODULE_PATH $OUT_PATH/orm

#替换go文件中的AAAAAAA 为 0
find $OUT_PATH/orm -type f |xargs -i sed -i s#AAAAAAA#0# {}

#编译go
cd main
GOOS=linux GOARCH=amd64 go build -o fantasy-db