package net

import (
	"fmt"
	"github.com/zhangyile1991911/cherry/iface"
)

type MessageHandler struct{
	APIS map[uint32] iface.IRouter
}

func NewMessageHandler()*MessageHandler{
	return &MessageHandler{APIS:make(map[uint32]iface.IRouter)}
}

func (this *MessageHandler)DispatchMsgHandler(request iface.IRequest){
	if h, ok := this.APIS[request.GetMsgID()];ok{
		h.PreHandle(request)
		h.Handle(request)
		h.PostHandle(request)
	}else{
		fmt.Printf("api msgid = %d missing\n",request.GetMsgID())
	}
}

func  (this *MessageHandler)AddRouter(msgID uint32,router iface.IRouter){
	if _,ok := this.APIS[msgID];!ok{
		this.APIS[msgID] = router
	}
}