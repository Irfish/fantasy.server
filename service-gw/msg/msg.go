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
}
