#!/usr/bin/env bash
#全局变量
SSH_ADDR="root@192.168.0.131"
ROOT_PATH=`pwd`
SERVER_NAME="fantasy-s"
PUBLISH_PATH="$ROOT_PATH/publish"
CONFIG_PATH="$ROOT_PATH/config"
SERVER_ROOT_PATH="/usr/fantasy/s"

#编译go
echo "build go start..."
GOOS=linux GOARCH=amd64 go build -o $SERVER_NAME
echo "build go end"
echo ""

echo "clear publish file start..."
cd $PUBLISH_PATH
pwd
rm -f -r *
cd -
echo "clear publish file end"
echo ""

echo "move files to publish start..."
mv $SERVER_NAME $PUBLISH_PATH
cp -r $CONFIG_PATH $PUBLISH_PATH
echo "move files to publish end"
echo ""

cd $PUBLISH_PATH
TARGET_TGZ=publish-`date +%Y-%m-%d_%H-%M-%S`.tgz
echo "tar -czvf "$TARGET_TGZ" start..."
tar -czvf $TARGET_TGZ *
md5sum $TARGET_TGZ
echo "tar -czvf end"
echo ""

echo "scp start..."
scp $TARGET_TGZ $SSH_ADDR:$SERVER_ROOT_PATH
if [ $? -ne 0 ]
then
    echo "scp fail"
    cd -
    exit 1
fi
#rm -f $TARGET_TGZ
echo "scp end"
echo ""

echo "ssh for tar -zxvf start..."
ssh $SSH_ADDR "tar -zxvf $SERVER_ROOT_PATH/"$TARGET_TGZ" -C  $SERVER_ROOT_PATH/"
if [ $? -ne 0 ]
then
    echo "ssh for tar -zxvf fail"
    cd -
    exit 1
fi
echo "ssh for tar -zxvf end"
echo ""

#给文件设置可执行权限
echo "ssh for chmod start..."
ssh $SSH_ADDR "chmod 777  $SERVER_ROOT_PATH/$SERVER_NAME"
if [ $? -ne 0 ]
then
    echo "chmod fail"
    cd -
    exit 1
fi
echo "ssh for chmod end"
cd -


echo "end"
