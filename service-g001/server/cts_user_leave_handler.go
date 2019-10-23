package server

import (
	"github.com/Irfish/component/leaf/gate"
	"github.com/Irfish/component/log"
	"github.com/Irfish/fantasy.server/pb"
	"github.com/Irfish/fantasy.server/service-g001/logic"
)

func ctsUserLeaveHandler(args []interface{}) {
	m := args[0].(*pb.CtsUserLeave)
	log.Debug("user leave:%s", m.UserId)
	p := logic.RoomManager.PlayerLeaveRoom(m.UserId)
	a := args[1].(gate.Agent)
	sendMessage(a, &pb.StcUserLeave{
		UserId:  p.UserId,
		RoomId:  p.RoomId,
		ChairId: p.ChairId,
	})
}
