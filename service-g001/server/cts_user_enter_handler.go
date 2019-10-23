package server

import (
	"github.com/Irfish/component/leaf/gate"
	"github.com/Irfish/component/log"
	"github.com/Irfish/fantasy.server/pb"
	"github.com/Irfish/fantasy.server/service-g001/logic"
)

func ctsUserEnterHandler(args []interface{}) {
	m := args[0].(*pb.CtsUserEnter)
	log.Debug("user enter:%s", m.UserId)
	p := logic.RoomManager.PlayerEnterRoom(m.UserId)
	a := args[1].(gate.Agent)
	sendMessage(a, &pb.StcUserEnter{
		UserId:  p.UserId,
		RoomId:  p.RoomId,
		ChairId: p.ChairId,
	})
}
