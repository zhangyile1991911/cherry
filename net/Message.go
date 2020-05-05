package net

type Message struct{
	Id uint32
	DateLen uint32
	Data []byte
}

func(m *Message)GetMsgId() uint32{
	return m.Id
}

func(m *Message)GetMsgLen() uint32{
	return m.DateLen
}

func(m *Message)GateData() []byte{
	return m.Data
}

func(m *Message)SetMsgId(id uint32){
	m.Id = id
}

func(m *Message)SetData(data []byte){
	m.Data = data
}

func(m *Message)SetMsgLen(len uint32){
	m.DateLen = len
}