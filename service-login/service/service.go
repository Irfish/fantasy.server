package service

import (
	"github.com/Irfish/component/etcd3"
	"github.com/Irfish/component/log"
	"github.com/Irfish/component/redis"
	"github.com/Irfish/component/xorm"
	"github.com/Irfish/fantasy.server/service-login/base"
	mGin "github.com/Irfish/fantasy.server/service-login/gin"
	"github.com/Irfish/fantasy.server/service-login/server"
	"os"
	"os/signal"
	"syscall"
)

func Run() {
	log.Debug("fantasy service login running ")
	etcd3.Init([]string{base.Server.EtcdAddr}, 3)
	redis.Run()
	xorm.Run()
	mGin.Run()
	server.Run()
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	<-ch
}
