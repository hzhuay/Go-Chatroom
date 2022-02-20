package main

import (
	"chatroom/server/model"
	"fmt"
	"net"
	"time"
)

func process(conn net.Conn) {
	defer conn.Close()
	//创建总控
	processor := &Processor{
		Conn: conn,
	}
	err := processor.process2()
	if err != nil {
		fmt.Println("客户端通讯协程出错", err)
		return
	}
}

func initUserDao() {
	model.MyUserDao = model.NewUserDao(pool)
}

//自动调用初始化函数
func init() {
	//初始化连接池
	initPool("112.124.37.136:6379", 16, 0, 300*time.Second)
	initUserDao()
}

func main() {
	fmt.Println("服务器在8889端口监听...")
	listen, err := net.Listen("tcp", "0.0.0.0:8889")
	if err != nil {
		fmt.Println("net.Listen err = ", err)
		return
	}
	defer listen.Close()

	for {
		fmt.Println("等待客户端连接服务器...")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("listen.Accept err = ", err)
		}
		//连接成功则启动一个协程和客户端保持连接
		go process(conn)
	}
}
