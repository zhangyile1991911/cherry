package iface

type IMessage interface {

	GetMsgId() uint32
	GetMsgLen() uint32
	GetData() []byte

	SetMsgId(uint32)
	SetData([]byte)
	SetMsgLen(uint32)
}
