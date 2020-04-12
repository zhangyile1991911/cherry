package net

import "github.com/zhangyile1991911/cherry/iface"

type BaseRoute struct{

}

func(b *BaseRoute) PreHandle(request iface.IRequest){

}

func(b *BaseRoute)Handle(request iface.IRequest){

}

func(b *BaseRoute)PostHandle(request iface.IRequest){

}