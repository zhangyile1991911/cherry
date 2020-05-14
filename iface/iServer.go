package iface

type IServer interface {
	Start()

	Stop()

	Run()

	AddRouter(msgId uint32,router IRouter)
}