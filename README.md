# fantasy.server (游戏服务器)
基于leaf编写的分布式服务器框架，其中对leaf做了一些修改和扩展，并上传到component中，服务器之间使用rpc或tcp进行交互

1. 使用etcd做服务发现

2. 使用redis作缓存

3. mysql做持久化存储

#service-db
处理持久化数据 
#service-game
游戏逻辑服
#service-gw
网关服务器,负责转发客户端消息
#service-log
处理日志
#service-login
登陆
#service-web
web服务器
#b.sh
编译项目并上传到远程服务器
#运行
每个服务可单独运行，也可以通过./b.sh 生成的fantasy-s运行
    运行gateway服务器
    ./fantasy-s -s gw
          


        



