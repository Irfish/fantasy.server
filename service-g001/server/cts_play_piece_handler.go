package server

import (
	"github.com/Irfish/component/leaf/gate"
	"github.com/Irfish/fantasy.server/pb"
	"github.com/Irfish/fantasy.server/service-g001/logic"
)

func ctsPlayPieceHandler(args []interface{}) {
	m := args[0].(*pb.CtsPlayPiece)
	a := args[1].(gate.Agent)
	//	log.Debug("play piece:%d %d %d %d", m.RoomId, m.ChairId, m.X, m.Y)
	list, e := logic.RoomManager.PlayPiece(m.ChairId, m.RoomId, m.X, m.Y)
	if e != nil {
		sendMessage(a, &pb.StcErrorNotice{
			Info: e.Error(),
		})
		return
	}
	sendMessage(a, &pb.StcPlayPiece{
		X:       m.X,
		Y:       m.Y,
		ChairId: m.ChairId,
		RoomId:  m.RoomId,
	})
	if len(list) >= logic.WinCount {
		//log.Debug("win:%v", list)
		sendMessage(a, &pb.StcGameResult{
			PieceList: list,
		})
	}
}
