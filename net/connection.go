package net

import (
	"fmt"
	"github.com/zhangyile1991911/cherry/iface"
	"net"
)

type Connection struct {
	Conn      *net.TCPConn
	ConnID    uint32
	isClose   bool
	handleAPI iface.HandleFunc
	ExitChan  chan bool
}

func NewConnection(conn *net.TCPConn, connID uint32, callback iface.HandleFunc) *Connection {
	c := new(Connection)
	c.Conn = conn
	c.ConnID = connID
	c.handleAPI = callback
	c.isClose = false
	c.ExitChan = make(chan bool, 1)

	return c
}

func (c *Connection)StartReader(){
	fmt.Println("Reader Goroutine is running...")
	defer fmt.Printf("connID = %d Reader is exit remote addr is %s",c.ConnID,c.GetRemoteAddr().String())
	defer c.Stop()

	buf := make([]byte,512)
	for{

		cnt,err := c.Conn.Read(buf)
		if err != nil{
			fmt.Printf("recv buf err %v",err)
			break
		}

		if c.handleAPI != nil {
			if err := c.handleAPI(c.Conn,buf,cnt);err != nil {
				fmt.Printf("ConnID %d handle is error %v",c.ConnID,err)
				break
			}
		}

	}

}

func (c *Connection) Start() {
	fmt.Println("Conn Start().. ConnID = ",c.ConnID)

	go c.StartReader()

}

func (c *Connection) Stop() {
	fmt.Println("Conn Stop().. ConnID = ",c.ConnID)
	if c.isClose {
		return
	}

	_ = c.Conn.Close()
	close(c.ExitChan)

	c.isClose = true
}

func (c *Connection) GetTCPConnection() *net.TCPConn {
	return c.Conn
}

func (c *Connection) GetConnID() uint32 {
	return c.ConnID
}

func (c *Connection) GetRemoteAddr() net.Addr{
	return c.Conn.RemoteAddr()
}

func (c *Connection) Send(data []byte) error {
	return nil
}
