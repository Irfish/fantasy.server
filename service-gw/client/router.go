package client

import (
	"reflect"

	"github.com/Irfish/fantasy.server/pb"
	"github.com/Irfish/fantasy.server/service-gw/msg"
)

func (c *Client) handler(m interface{}, h interface{}) {
	c.AgentChanRPC.Register(reflect.TypeOf(m), h)
}

func (c *Client) initHandler() {
	c.router()
	c.registerHandler()
}

func (c *Client) router() {
	msg.Processor.SetRouter(&pb.StcUserEnter{}, c.AgentChanRPC)
}

func (c *Client) registerHandler() {
	c.handler(&pb.StcUserEnter{}, stcUserEnterHandler)
}
