package net

import "github.com/zhangyile1991911/cherry/iface"

type Request struct{
	conn iface.IConnection
	msg iface.IMessage
}

func (r *Request) GetConnection() iface.IConnection{
	return r.conn
}

func (r *Request)GetData()[]byte{
	return r.msg.GetData()
}

func (r *Request)GetMsgID()uint32{
	return r.msg.GetMsgId()
}