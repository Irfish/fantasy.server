package service

import (
	"reflect"

	"github.com/Irfish/component/etcd3"
	"github.com/Irfish/component/leaf"
	lconf "github.com/Irfish/component/leaf/conf"
	"github.com/Irfish/component/log"
	"github.com/Irfish/component/redis"
	"github.com/Irfish/fantasy.server/service-gw/base"
	"github.com/Irfish/fantasy.server/service-gw/msg"
	"github.com/Irfish/fantasy.server/service-gw/server"
)

func Run() {
	log.Debug("fantasy service gw running ")
	lconf.LogLevel = base.Server.LogLevel
	lconf.LogPath = base.Server.LogPath
	lconf.LogFlag = base.LogFlag
	lconf.ConsolePort = base.Server.ConsolePort
	lconf.ProfilePath = base.Server.ProfilePath
	redis.Run()
	//连接etcd
	etcd3.Init([]string{base.Server.EtcdAddr}, 3)
	msg.Processor.Range(func(id uint16, t reflect.Type) {
		log.Debug("message server: id =%d,%s", id, t.String())
	})
	msg.Processor.Range(func(id uint16, t reflect.Type) {
		log.Debug("message client: id =%d,%s", id, t.String())
	})
	leaf.Run(server.Module)
}
