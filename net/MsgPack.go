package net

import (
	"bytes"
	"encoding/binary"
	"errors"
	"github.com/zhangyile1991911/cherry/iface"
	"github.com/zhangyile1991911/cherry/utilis"
)


func GetHeadLen() uint32{
	//len uint32(4 byte) + id uint32(4 byte)
	return 8
}

func Pack(msg iface.IMessage)([]byte,error){
	dataBuff := bytes.NewBuffer([]byte{})
	if err := binary.Write(dataBuff,binary.LittleEndian,msg.GetMsgLen());err != nil{
		return nil,err
	}
	if err := binary.Write(dataBuff,binary.LittleEndian,msg.GetMsgId());err != nil{
		return nil,err
	}
	if err := binary.Write(dataBuff,binary.LittleEndian,msg.GetData());err != nil{
		return nil,err
	}
	return dataBuff.Bytes(),nil
}

func Unpack(byteData []byte)(iface.IMessage,error){
	dataBuff := bytes.NewReader(byteData)


	msg := &Message{}

	if err := binary.Read(dataBuff,binary.LittleEndian,&msg.DataLen);err != nil{
		return nil,err
	}

	if msg.DataLen > utilis.GlobalObj.MaxPackageSize{
		return nil,errors.New("msg data is too large")
	}

	if err := binary.Read(dataBuff,binary.LittleEndian,&msg.Id);err != nil{
		return nil,err
	}

	//if err := binary.Read(dataBuff,binary.LittleEndian,&msg.Data);err != nil{
	//	return nil,err
	//}

	return msg,nil
}