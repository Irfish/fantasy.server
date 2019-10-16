package base

import (
	"encoding/json"
	"io/ioutil"
	"time"

	log1 "log"

	"github.com/Irfish/component/log"
)

var (
	// log config
	LogFlag = log1.LstdFlags
	// gate config
	PendingWriteNum        = 2000
	MaxMsgLen       uint32 = 4096
	HTTPTimeout            = 10 * time.Second
	LenMsgLen              = 2
	LittleEndian           = false
	// skeleton config
	GoLen              = 10000
	TimerDispatcherLen = 10000
	AsynCallLen        = 10000
	ChanRPCLen         = 10000
)

var Server struct {
	WSAddr      string
	TCPAddr     string
	LogLevel    string
	LogPath     string
	CertFile    string
	KeyFile     string
	MaxConnNum  int
	ConsolePort int
	ProfilePath string
	EtcdAddr    string
}

func init() {
	data, err := ioutil.ReadFile("config/gw.json")
	if err != nil {
		log.Fatal("%v", err)
	}
	err = json.Unmarshal(data, &Server)
	if err != nil {
		log.Fatal("%v", err)
	}
}
