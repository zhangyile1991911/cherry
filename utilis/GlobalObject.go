package utilis

import (
	"encoding/json"
	"github.com/zhangyile1991911/cherry/iface"
	"io/ioutil"
)

type GlobalObject struct{
	TCPServer iface.IServer
	Host string `json:"host"`
	Port int `json:"port"`
	Name string `json:"name"`

	Version string `json:"Version"`
	MaxConn int `json:"max_conn"`
	MaxPackageSize uint32 `json:"max_package_size"`
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
