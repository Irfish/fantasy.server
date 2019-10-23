package logic

import (
	"github.com/Irfish/fantasy.server/service-g001/base"
)

func (r *Room) OnInit() {
	r.Skeleton = base.NewSkeleton()
	r.Skeleton.TimerDispatcherLen = 1
}
func (r *Room) BeforeDestroy() {

}
func (r *Room) OnDestroy() {

}

func (r *Room) Run(closeSig chan bool) {
	go r.Skeleton.Run(closeSig)
}


