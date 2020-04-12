package net

import (
	"fmt"
	"github.com/zhangyile1991911/cherry/iface"
	"net"
)

type TCPServer struct {
	Name    string
	Network string
	Addr    string
	Port    int
	Router	iface.IRouter
}

func (s *TCPServer) Start() {
	fmt.Printf("[Start] Server Listenner at IP:%s,Port %d is starting\n", s.Addr, s.Port)
	tcpAddr, err := net.ResolveTCPAddr(s.Network, s.Addr)
	if err != nil {
		fmt.Println("resolve tcp addr error ", err)
		return
	}

	listener, err := net.ListenTCP(s.Network, tcpAddr)
	if err != nil {
		fmt.Println("ListenTCP error ", err)
		return
	}

	var cid uint32
	cid = 0
	fmt.Printf("start cherry server success %s Listenning\n", s.Name)
	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			fmt.Printf("Accept err", err)
		}
		dealConn := NewConnection(conn, cid, s.Router)
		dealConn.Start()
		cid++
	}
}

//func EchoClient(conn *net.TCPConn, data []byte, cnt int) error {
//	fmt.Printf("[Conn Hanlde] CallBackToClient\n")
//	if _, err := conn.Write(data[:cnt]); err != nil {
//		return errors.New("Echo error ")
//	}
//	return nil
//}

func (s *TCPServer) Stop() {

}

func (s *TCPServer) Run() {
	s.Start()
}

func (s *TCPServer)AddRouter(router iface.IRouter){
	s.Router = router
}

func NewServer(name string) iface.IServer {
	s := new(TCPServer)
	s.Name = name
	s.Network = "tcp4"
	s.Addr = "0.0.0.0:8999"
	s.Port = 8999
	return s
}
