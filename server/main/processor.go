package main

import (
	"chatroom/common/message"
	"chatroom/server/processes"
	"chatroom/server/utils"
	"fmt"
	"io"
	"net"
)

type Processor struct {
	Conn net.Conn
}

func (this *Processor) serverProcessMes(mes *message.Message) (err error) {
	switch mes.Type {
	case message.LoginMesType:
		//处理登录请求
		//创建一个UserProcess实例
		up := &processes.UserProcess{
			Conn: this.Conn,
		}
		err = up.ServerProcessLogin(mes)
	case message.RegisterMesType:
		//处理注册
	default:
		fmt.Println("消息类型不存在！")
	}
	return
}

func (this *Processor) process2() (err error) {
	for {
		tf := &utils.Transfer{
			Conn: this.Conn,
		}

		mes, err := tf.ReadPkg()
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端断开了连接")
			} else {
				fmt.Println("ReadPkg err = ", err)
			}
			return err
		}
		fmt.Println("mes = ", mes)

		err = this.serverProcessMes(&mes)
		if err != nil {
			fmt.Println("process 2", err)
			return err
		}
	}
}
