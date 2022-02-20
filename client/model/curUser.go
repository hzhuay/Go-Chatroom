package model

import (
	"chatroom/common/message"
	"net"
)

//在客户端很多地方都会用到这个，所以要作为全局变量
type CurUser struct {
	Conn net.Conn
	message.User
}
