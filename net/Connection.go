package net

import (
	"fmt"
	"github.com/zhangyile1991911/cherry/iface"
	"net"
)

type IConnection struct {
	Conn      *net.TCPConn
	ConnID    uint32
	isClose   bool
	//handleAPI iface.HandleFunc
	ExitChan  chan bool
	Router    iface.IRouter
}

func NewConnection(conn *net.TCPConn, connID uint32, router iface.IRouter) *IConnection {
	c := new(IConnection)
	c.Conn = conn
	c.ConnID = connID
	c.Router = router
	c.isClose = false
	c.ExitChan = make(chan bool, 1)

	return c
}

func (c *IConnection) StartReceive() {
	fmt.Println("Reader Goroutine is running...")
	defer fmt.Printf("connID = %d Reader is exit remote addr is %s", c.ConnID, c.GetRemoteAddr().String())
	defer c.Stop()

	buf := make([]byte, 512)
	for {

		_, err := c.Conn.Read(buf)
		if err != nil {
			fmt.Printf("recv buf err %v", err)
			break
		}

		req := Request{conn:c,data:buf}

		c.Router.PreHandle(&req)

		c.Router.Handle(&req)

		c.Router.PostHandle(&req)
	}

}

func (c *IConnection) Start() {
	fmt.Println("Conn Start().. ConnID = ", c.ConnID)

	go c.StartReceive()

}

func (c *IConnection) Stop() {
	fmt.Println("Conn Stop().. ConnID = ", c.ConnID)
	if c.isClose {
		return
	}

	_ = c.Conn.Close()
	close(c.ExitChan)

	c.isClose = true
}

func (c *IConnection) GetTCPConnection() *net.TCPConn {
	return c.Conn
}

func (c *IConnection) GetConnID() uint32 {
	return c.ConnID
}

func (c *IConnection) GetRemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

func (c *IConnection) Send(data []byte) error {
	return nil
}
