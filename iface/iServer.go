package iface

type IServer interface {
	Start()

	Stop()

	Run()

	AddRouter(router IRouter)
}