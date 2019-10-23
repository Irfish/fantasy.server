package logic

const (
	PLAYER_STATUS_ONLINE = iota
	PLAYER_STATUS_OFFLINE
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

func (p *Player) Play(x, y int) {
	piece := Piece{Value: p.PieceColor, X: x, Y: y}
	PlayPiece(piece)
}
