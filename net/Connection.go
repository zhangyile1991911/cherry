package net

import (
	"errors"
	"fmt"
	"github.com/zhangyile1991911/cherry/iface"
	"io"
	"net"
)

type Connection struct {
	Conn      *net.TCPConn
	ConnID    uint32
	isClose   bool
	//handleAPI iface.HandleFunc
	ExitChan  chan bool
	Router    iface.IRouter
}

func NewConnection(conn *net.TCPConn, connID uint32, router iface.IRouter) *Connection {
	c := new(Connection)
	c.Conn = conn
	c.ConnID = connID
	c.Router = router
	c.isClose = false
	c.ExitChan = make(chan bool, 1)

	return c
}

func (c *Connection) StartReceive() {
	fmt.Println("Reader Goroutine is running...")
	defer fmt.Printf("connID = %d Reader is exit remote addr is %s", c.ConnID, c.GetRemoteAddr().String())
	defer c.Stop()

	for {
		headData := make([]byte,GetHeadLen())
		if _,err := io.ReadFull(c.Conn,headData);err != nil{
			fmt.Println("read msg head error",err)
			break
		}

		msg,err := Unpack(headData)
		if err != nil{
			fmt.Println("")
			break
		}

		if msg.GetMsgLen() <= 0 {
			fmt.Println("")
			break
		}

		bodyData := make([]byte,msg.GetMsgLen())
		if _,err := io.ReadFull(c.Conn,bodyData);err != nil{
			fmt.Println("")
			break
		}

		msg.SetData(bodyData)

		req := &Request{
			conn:c,
			msg:msg}

		c.Router.PreHandle(req)

		c.Router.Handle(req)

		c.Router.PostHandle(req)
	}

}


func (c *Connection)SendMsg(msgId uint32,data []byte)error{
	if c.isClose{
		return errors.New("")
	}

	msg := &Message{}
	msg.Id = msgId
	msg.DataLen = uint32(len(data))
	msg.Data = data


	byteDatas,err := Pack(msg)
	if err != nil{
		fmt.Println("")
		return err
	}

	if _,err := c.Conn.Write(byteDatas);err != nil{
		fmt.Println("")
		return err
	}

	return nil
}

func (c *Connection) Start() {
	fmt.Println("Conn Start().. ConnID = ", c.ConnID)

	go c.StartReceive()

}

func (c *Connection) Stop() {
	fmt.Println("Conn Stop().. ConnID = ", c.ConnID)
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

func (c *Connection) GetRemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

