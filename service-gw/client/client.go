package client

import (
	"github.com/Irfish/component/leaf/chanrpc"
	"github.com/Irfish/component/leaf/tcpclient"
	"github.com/Irfish/component/log"
	"github.com/Irfish/fantasy.server/pb"
	"github.com/Irfish/fantasy.server/service-gw/base"
	"github.com/Irfish/fantasy.server/service-gw/msg"
	"github.com/micro/protobuf/proto"
)

var (
	serviceToAgent = make(map[string]tcpclient.Agent, 0)
)

func NewClient(addr string, id string, autoReconnect bool) *Client {
	c := new(Client)
	c.Id = id
	c.ServiceAddr = addr
	c.AutoReconnect = autoReconnect
	c.Init()
	return c
}

type Client struct {
	*tcpclient.Client
	Id            string
	ServiceAddr   string
	AutoReconnect bool
	Agent         tcpclient.Agent
	ServerChanRpc *chanrpc.Server
}

func (c *Client) AppendOuterChanRpc(name string, chanRpc *chanrpc.Server) {
	c.Client.OuterChanRPC[name] = chanRpc
}

func (c *Client) Init() {
	skeleton := base.NewSkeleton()
	c.Client = &tcpclient.Client{
		PendingWriteNum: base.PendingWriteNum,
		MaxMsgLen:       base.MaxMsgLen,
		HTTPTimeout:     base.HTTPTimeout,
		TCPAddr:         c.ServiceAddr,
		LenMsgLen:       base.LenMsgLen,
		LittleEndian:    true,
		Processor:       msg.Processor,
		AgentChanRPC:    skeleton.ChanRPCServer,
		Id:              c.Id,
		AutoReconnect:   c.AutoReconnect,
		Skeleton:        skeleton,
		OuterChanRPC:    make(map[string]*chanrpc.Server, 0),
	}
	c.InitRegister()
	c.initHandler()
}

func (c *Client) SendToService(msg []byte,id int32) {
	if a, ok := serviceToAgent[c.Id]; ok {
		a.WriteMsg(&pb.Message{Body: msg, Header: &pb.Header{}})
		return
	}
	log.Debug("client agent not found:%s", c.Id)
}

func sendMessage(a tcpclient.Agent, m interface{}) {
	m1 := m.(proto.Message)
	body, err := msg.Processor.Marshal(m1)
	if err != nil {
		log.Error("SendToService proto.Marshal message err:%s", err.Error())
	}
	bytes := make([]byte, 0)
	for _, b := range body {
		bytes = append(bytes, b...)
	}
	a.WriteMsg(&pb.Message{Body: bytes, Header: &pb.Header{}})
}
