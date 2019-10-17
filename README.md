#### fantasy.server (游戏服务器)
基于[leaf](https://github.com/name5566/leaf)编写的分布式服务器框架，其中对leaf做了一些修改和扩展，并上传到[component](https://github.com/Irfish/component)中，服务器之间使用rpc或tcp进行交互

1. 使用etcd做服务发现

2. 使用redis作缓存

3. mysql做持久化存储

#### service-db
处理持久化数据
#### service-game
游戏逻辑服
#### service-gw
网关服务器,负责转发客户端消息
#### service-log
处理日志
#### service-login
登陆
#### service-web
web服务器

##### p.sh
编译项目并上传到远程服务器


##### 运行服务:

使用./b_linux.sh  生成二进制可执行文件 fantasy-s

    ./fantasy-s -s gw  (linux)
    
   or
   
 使用./b_win.sh 生成exe可执行文件 fantasy-s.exe
 
    ./fantasy-s.exe -s g001 (win)
    

每个服务也可以单独编译并运行


          


        



