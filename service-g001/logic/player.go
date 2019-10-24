package logic

import "github.com/Irfish/fantasy.server/pb"

const (
	PlayerStatusOnline = iota
	PlayerStatusOffline
)

type Player struct {
	UserId     int64
	Gold       int64
	Name       string
	Status     int32
	PieceColor int32
	RoomId     int64
	ChairId    int32
}

func (p *Player) Play(x, y int32) (list []*pb.Piece, e error) {
	piece := pb.Piece{Value: p.PieceColor, X: x, Y: y}
	list, e = PlayPiece(piece)
	return
}
