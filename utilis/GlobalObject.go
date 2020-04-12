package utilis

import (
	"encoding/json"
	"github.com/zhangyile1991911/cherry/iface"
	"io/ioutil"
)

type GlobalObject struct{
	TCPServer iface.IServer
	Host string
	Port int
	Name string

	Version string
	MaxConn int
	MaxPackageSize uint32
}

var GlobalObj *GlobalObject

func (g *GlobalObject)Reload(){
	data,err := ioutil.ReadFile("config/cherry.json")
	if err != nil{
		panic(err)
	}
	err = json.Unmarshal(data,g)
	if err != nil{
		panic(err)
	}
}

func init(){
	GlobalObj = new(GlobalObject)
	GlobalObj.Name = ""
	GlobalObj.Version = ""
	GlobalObj.Port = 9998
	GlobalObj.Host = "0.0.0.0"
	GlobalObj.MaxConn = 1000
	GlobalObj.MaxPackageSize = 4096
	GlobalObj.Reload()
}
