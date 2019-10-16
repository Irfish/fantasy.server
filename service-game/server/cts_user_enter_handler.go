package server

import (
	"github.com/Irfish/component/leaf/gate"
	"github.com/Irfish/component/log"
	"github.com/Irfish/fantasy.server/pb"
)

func ctsUserEnterHandler(args []interface{}) {
	m := args[0].(*pb.CtsUserEnter)
	log.Debug("received message from client:%s", m.Token)
	a := args[1].(gate.Agent)
	sendMessage(a, &pb.StcUserEnter{Result: "succeed enter"})
}
