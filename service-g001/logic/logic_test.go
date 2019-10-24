package logic

import (
	"fmt"
	"testing"

	"github.com/Irfish/fantasy.server/pb"
)

func TestPlayPiece(t *testing.T) {
	InitTable(10,10)
	p1:=pb.Piece{
		Value:1,
		X:0,
		Y:0,
	}
	_,e:= PlayPiece(p1)
	if e!=nil{
		fmt.Println(e.Error())
		return
	}
}
