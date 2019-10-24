package msg

import (
	"github.com/Irfish/component/leaf/network/protobuf"
	"github.com/Irfish/fantasy.server/pb"
)

var Processor = protobuf.NewProcessor()

func init() {
	Processor.SetByteOrder(true)
	Processor.Register(&pb.Message{})
	Processor.Register(&pb.StcErrorNotice{})
	Processor.Register(&pb.CtsUserAuthentication{})
	Processor.Register(&pb.StcUserAuthentication{})
	Processor.Register(&pb.CtsUserEnter{})
	Processor.Register(&pb.StcUserEnter{})
	Processor.Register(&pb.CtsUserLeave{})
	Processor.Register(&pb.StcUserLeave{})
	Processor.Register(&pb.CtsCreateRoom{})
	Processor.Register(&pb.StcCreateRoom{})
	Processor.Register(&pb.CtsPlayPiece{})
	Processor.Register(&pb.StcPlayPiece{})
	Processor.Register(&pb.StcGameResult{})
	Processor.Register(&pb.CtsUserReady{})
	Processor.Register(&pb.StcUserReady{})
}
