package logic

import (
	"fmt"

	"github.com/Irfish/fantasy.server/pb"
)

var (
	TablePanel [][]int32
	MaxX       int32
	MaxY       int32
)

const (
	PieceValueZero  = 0 //空白棋子
	PieceValueBlack = 1 //黑棋
	PieceValueWhite = 2 //白棋

	WinCount = 5 //最大相连数
	//检测方向
	DirectionHorizontal = 1 //水平
	DirectionVertical   = 2 //垂直
	DirectionLeftTilt   = 3 //左倾斜
	DirectionRightTilt  = 4 //右倾斜
)

func InitTable(maxX, maxY int) {
	MaxX = int32(maxX)
	MaxY = int32(maxY)
	TablePanel = make([][]int32, maxX)

	for i := 0; i < maxX; i++ {
		TablePanel[i] = make([]int32, maxY)
		for j := 0; j < maxY; j++ {
			TablePanel[i][j] = PieceValueZero
		}
	}
}

func PlayPiece(piece pb.Piece) (list []*pb.Piece, e error) {
	if piece.X >= MaxX || piece.Y >= MaxY || piece.X < 0 || piece.Y < 0 {
		e = fmt.Errorf("x or y not illegal")
		return
	}
	if TablePanel[piece.X][piece.Y] == 0 {
		TablePanel[piece.X][piece.Y] = piece.Value
	}
	list = CheckWin(piece)
	return
}

//以落子的点为中心进行查找
func CheckWin(piece pb.Piece) (list []*pb.Piece) {
	//四个方向 1,2,3,4
	p1 := check(piece, DirectionHorizontal)
	if len(p1) >= WinCount {
		list = append(list, p1...)
	}
	p2 := check(piece, DirectionVertical)
	if len(p2) >= WinCount {
		list = append(list, p2...)
	}
	p3 := check(piece, DirectionLeftTilt)
	if len(p3) >= WinCount {
		list = append(list, p3...)
	}
	p4 := check(piece, DirectionRightTilt)
	if len(p4) >= WinCount {
		list = append(list, p4...)
	}
	return
}

func check(piece pb.Piece, d int) (pieceList []*pb.Piece) {
	switch d {
	case DirectionHorizontal: //横向
		x := piece.X
		for {
			if TablePanel[x][piece.Y] == piece.Value {
				pieceList = append(pieceList, &pb.Piece{X: x, Y: piece.Y, Value: piece.Value})
			} else {
				break
			}
			x++
			if x >= MaxX {
				break
			}
		}
		x = piece.X - 1
		for {
			if x < 0 {
				break
			}
			if TablePanel[x][piece.Y] == piece.Value {
				pieceList = append(pieceList, &pb.Piece{X: x, Y: piece.Y, Value: piece.Value})
			} else {
				break
			}
			x--
		}
	case DirectionVertical: //纵向
		y := piece.Y
		for {
			if TablePanel[piece.X][y] == piece.Value {
				pieceList = append(pieceList, &pb.Piece{X: piece.X, Y: y, Value: piece.Value})
			} else {
				break
			}
			y++
			if y >= MaxY {
				break
			}
		}
		y = piece.Y - 1
		for {
			if y < 0 {
				break
			}
			if TablePanel[piece.X][y] == piece.Value {
				pieceList = append(pieceList, &pb.Piece{X: piece.X, Y: y, Value: piece.Value})
			} else {
				break
			}
			y--
		}
	case DirectionLeftTilt:
		x := piece.X
		y := piece.Y
		for {
			if TablePanel[x][y] == piece.Value {
				pieceList = append(pieceList, &pb.Piece{X: x, Y: y, Value: piece.Value})
			} else {
				break
			}
			x++
			y++
			if x >= MaxX || y >= MaxY {
				break
			}
		}
		x = piece.X - 1
		y = piece.Y - 1
		for {
			if x < 0 || y < 0 {
				break
			}
			if TablePanel[x][y] == piece.Value {
				pieceList = append(pieceList, &pb.Piece{X: x, Y: y, Value: piece.Value})
			} else {
				break
			}
			x--
			y--
		}
	case DirectionRightTilt:
		x := piece.X
		y := piece.Y
		for {
			if TablePanel[x][y] == piece.Value {
				pieceList = append(pieceList, &pb.Piece{X: x, Y: y, Value: piece.Value})
			} else {
				break
			}
			x--
			y++
			if x < 0 || y >= MaxY {
				break
			}
		}
		x = piece.X + 1
		y = piece.Y - 1
		for {
			if x >= MaxX || y < 0 {
				break
			}
			if TablePanel[x][y] == piece.Value {
				pieceList = append(pieceList, &pb.Piece{X: x, Y: y, Value: piece.Value})
			} else {
				break
			}
			x++
			y--
		}
	}
	return
}
