package iface


type IMessageHandler interface{
	DispatchMsgHandler(request IRequest)

	AddRouter(msgID uint32,router IRouter)
}
