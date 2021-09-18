package processes

import (
	"chatroom/client/utils"
	"chatroom/common/message"
	"encoding/json"
	"fmt"
	"net"
)

type UserProcess struct {
}

func (this *UserProcess) Login(userID int, userPwd string) (err error) {
	//连接到服务器
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("net.Dial err = ", err)
		return
	}
	defer conn.Close()

	//准备把消息发给服务器
	var mes message.Message
	mes.Type = message.LoginMesType
	var loginMes message.LoginMes
	loginMes.UserId = userID
	loginMes.UserPwd = userPwd

	//序列化
	data, err := json.Marshal(loginMes)
	if err != nil {
		fmt.Println("json.Marshal err = ", err)
		return
	}
	mes.Date = string(data)
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err = ", err)
		return
	}

	//处理服务器返回的消息
	tf := utils.Transfer{
		Conn: conn,
	}

	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("注册发送信息出错 ", err)
		return
	}

	// //发送包长度
	// var pkgLen uint32 = uint32(len(data))
	// var buf []byte = make([]byte, 4)
	// binary.BigEndian.PutUint32(buf, pkgLen)

	// _, err = conn.Write(buf)
	// if err != nil {
	// 	fmt.Println("conn.Write err = ", err)
	// 	return
	// }

	// //发送消息本身
	// _, err = conn.Write(data)
	// if err != nil {
	// 	fmt.Println("conn.Write(data) err = ", err)
	// 	return
	// }

	mes, err = tf.ReadPkg()
	if err != nil {
		fmt.Println("readPkg err =", err)
		return
	}
	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(mes.Date), &loginResMes)
	if loginResMes.Code == 200 {
		// fmt.Println("登录成功")
		//启动一个协程，保持 和服务器的连接，接收服务器的推送
		go serverProcessMes(conn)
		//循环显示登录成功后的菜单
		for {
			ShowMenu()
		}
	} else {
		fmt.Println(loginResMes.Error)
	}

	return
}

func (this *UserProcess) Register(userID int, userPwd string, userName string) (err error) {
	//连接到服务器
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("net.Dial err = ", err)
		return
	}
	defer conn.Close()

	//准备把消息发给服务器
	var mes message.Message
	mes.Type = message.RegisterMesType
	var registerMes message.RegisterMes
	registerMes.User.UserId = userID
	registerMes.User.UserPwd = userPwd
	registerMes.User.UserName = userName

	//序列化
	data, err := json.Marshal(registerMes)
	if err != nil {
		fmt.Println("json.Marshal err = ", err)
		return
	}
	mes.Date = string(data)
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err = ", err)
		return
	}

	//准备一个发送实例
	tf := utils.Transfer{
		Conn: conn,
	}

	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("注册发送信息出错 ", err)
		return
	}

	mes, err = tf.ReadPkg()
	if err != nil {
		fmt.Println("readPkg err =", err)
		return
	}

	//处理服务器返回信息
	var RegisterResMes message.RegisterResMes
	err = json.Unmarshal([]byte(mes.Date), &RegisterResMes)
	if RegisterResMes.Code == 200 {
		fmt.Println("注册成功，请重新登录")
		//启动一个协程，保持 和服务器的连接，接收服务器的推送
		// go serverProcessMes(conn)
		//循环显示登录成功后的菜单
		// for {
		// 	ShowMenu()
		// }
	} else {
		fmt.Println(RegisterResMes.Error)
	}
	return
}
