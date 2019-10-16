package base

import (
	"github.com/Irfish/component/leaf/chanrpc"
	"github.com/Irfish/component/leaf/module"
)

func NewSkeleton() *module.Skeleton {
	skeleton := &module.Skeleton{
		GoLen:              GoLen,
		TimerDispatcherLen: TimerDispatcherLen,
		AsynCallLen:        AsynCallLen,
		ChanRPCServer:      chanrpc.NewServer(ChanRPCLen),
	}
	skeleton.Init()
	return skeleton
}
