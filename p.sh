#!/usr/bin/env bash
#配置ssh免密登陆
#ssh-keygen   生成 私钥和公钥 or   ssh-keygen -t rsa
#cd ~/.ssh  查看秘钥是否生成（id_rsa:私钥  id_rsa.pub:公钥）
#ssh-copy-id -i ~/.ssh/id_rsa.pub root@remotehost 将公钥copy到远程服务器

#远程服务器地址
SSH_ADDR="root@remotehost"
#当前目录
ROOT_PATH=`pwd`
#可执行文件名称
SERVER_NAME="fantasy-s"
#发布文件临时存放目录
PUBLISH_PATH="$ROOT_PATH/publish"
#配置文件
CONFIG_PATH="$ROOT_PATH/config"
#文件远程存放目录
REMOTE_SERVER_ROOT_PATH="/usr/fantasy/s"

#编译二进制可执行文件
echo "build go start..."
GOOS=linux GOARCH=amd64 go build -o $SERVER_NAME
echo "build go end"
echo ""

#清空publish 目录
echo "clear publish file start..."
cd $PUBLISH_PATH
pwd
rm -f -r *
cd -
echo "clear publish file end"
echo ""

#将需要上传的文件放入publish目录中
echo "move files to publish start..."
mv $SERVER_NAME $PUBLISH_PATH
cp -r $CONFIG_PATH $PUBLISH_PATH
echo "move files to publish end"
echo ""

#打包压缩publish中的所有文件
cd $PUBLISH_PATH
TARGET_TGZ=publish-`date +%Y-%m-%d_%H-%M-%S`.tgz
echo "tar -czvf "$TARGET_TGZ" start..."
tar -czvf $TARGET_TGZ *
md5sum $TARGET_TGZ
echo "tar -czvf end"
echo ""

#上传文件到远程服务器
echo "scp start..."
scp $TARGET_TGZ $SSH_ADDR:$REMOTE_SERVER_ROOT_PATH
if [ $? -ne 0 ]
then
    echo "scp fail"
    cd -
    exit 1
fi
#rm -f $TARGET_TGZ
echo "scp end"
echo ""

#远程解压上传的文件
echo "ssh for tar -zxvf start..."
ssh $SSH_ADDR "tar -zxvf $REMOTE_SERVER_ROOT_PATH/"$TARGET_TGZ" -C  $REMOTE_SERVER_ROOT_PATH/"
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
ssh $SSH_ADDR "chmod 777  $REMOTE_SERVER_ROOT_PATH/$SERVER_NAME"
if [ $? -ne 0 ]
then
    echo "chmod fail"
    cd -
    exit 1
fi
echo "ssh for chmod end"
echo ""
cd -
echo "do end"
