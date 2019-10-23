package server

import (
	"github.com/Irfish/component/leaf/gate"
	"github.com/Irfish/component/log"
	"github.com/Irfish/fantasy.server/pb"
	"github.com/Irfish/fantasy.server/service-g001/logic"
)

func ctsCreateRoomHandler(args []interface{}) {
	m := args[0].(*pb.CtsCreateRoom)
	a := args[1].(gate.Agent)
	log.Debug("create room:%d", m.UserId)
	roomId, e := logic.RoomManager.CreateRoom(m.UserId)
	if e != nil {
		sendMessage(a, &pb.StcErrorNotice{
			Info: e.Error(),
		})
		return
	}
	sendMessage(a, &pb.StcCreateRoom{
		UserId: m.UserId,
		RoomId: roomId,
	})
}
