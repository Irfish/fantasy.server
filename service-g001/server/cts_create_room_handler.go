package server

import (
	"github.com/Irfish/component/leaf/gate"
	"github.com/Irfish/component/log"
	"github.com/Irfish/fantasy.server/pb"
	"github.com/Irfish/fantasy.server/service-g001/logic"
)

func ctsCreateRoomHandler(args []interface{}) {
	m := args[0].(*pb.CtsCreateRoom)
	log.Debug("create room:%s", m.UserId)
	p := logic.RoomManager.CreateRoom(m.UserId)
	a := args[1].(gate.Agent)
	sendMessage(a, &pb.StcCreateRoom{
		UserId:  p.UserId,
		RoomId:  p.RoomId,
		ChairId: p.ChairId,
	})
}
