package server

import (
	"github.com/Irfish/component/log"
	"github.com/Irfish/fantasy.server/pb"
)

func ctsUserEnterHandler(args []interface{}) {
	m := args[0].(*pb.CtsUserEnter)
	log.Debug("received message form client:%s", m.UserId)
}
