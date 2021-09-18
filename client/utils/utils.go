package utils

import (
	"chatroom/common/message"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"net"
)

//将这些方法关联到结构体中
type Transfer struct {
	Conn net.Conn
	Buf  [8096]byte
}

func (this *Transfer) ReadPkg() (mes message.Message, err error) {

	_, err = this.Conn.Read(this.Buf[:4])
	if err != nil {
		err = errors.New("read package header error")
		return
	}
	//把读到的包大小转换成uint32
	var pkgLen uint32 = binary.BigEndian.Uint32(this.Buf[:4])
	//根据包长度读取信息
	n, err := this.Conn.Read(this.Buf[:pkgLen])
	if n != int(pkgLen) || err != nil {
		// err = errors.New("read package body error")
		return
	}

	//把包反序列化成Mes，注意取地址
	err = json.Unmarshal(this.Buf[:pkgLen], &mes)
	if err != nil {
		err = errors.New("json Unmarshal error")
		return
	}
	return
}

func (this *Transfer) WritePkg(data []byte) (err error) {
	//先发送长度
	var pkgLen uint32 = uint32(len(data))
	binary.BigEndian.PutUint32(this.Buf[:4], pkgLen)

	n, err := this.Conn.Write(this.Buf[:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write err = ", err)
		return
	}

	//发送data本身
	n, err = this.Conn.Write(data)
	if n != int(pkgLen) || err != nil {
		fmt.Println("conn.Write(data) fail", err)
		return
	}
	return
}
