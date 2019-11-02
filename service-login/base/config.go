package base

import (
	"encoding/json"
	"github.com/Irfish/component/log"
	"io/ioutil"
)

var Server struct {
	LogLevel    string
	LogPath     string
	EtcdAddr    string
	GinAddr     string
}

func init() {
	data, err := ioutil.ReadFile("config/login.json")
	if err != nil {
		log.Fatal("%v", err)
	}
	err = json.Unmarshal(data, &Server)
	if err != nil {
		log.Fatal("%v", err)
	}
}
