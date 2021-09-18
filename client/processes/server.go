package processes

import (
	"chatroom/client/utils"
	"fmt"
	"net"
	"os"
)

//显示登录成功后的界面
func ShowMenu() {
	fmt.Println("----------恭喜XXX登录成功----------")
	fmt.Println("----------1. 显示在线用户列表----------")
	fmt.Println("----------2. 发送消息----------")
	fmt.Println("----------3. 消息列表----------")
	fmt.Println("----------4. 退出系统----------")
	fmt.Print("请选择(1-4): ")
	var key int
	fmt.Scanf("%d\n", &key)
	switch key {
	case 1:
		fmt.Println("----------1. 显示在线用户列表----------")
	case 2:
		fmt.Println("----------2. 发送消息----------")
	case 3:
		fmt.Println("----------3. 消息列表----------")
	case 4:
		fmt.Println("Bye bye!")
		os.Exit(0)
	default:
		fmt.Println("输入的指令不正确!")
	}

}

//和服务器保持通讯
func serverProcessMes(conn net.Conn) {
	//创建Transfer实例，不停地读取服务器信息
	tf := &utils.Transfer{
		Conn: conn,
	}
	for {
		// fmt.Println("客户端正在等待读取服务器发送的消息。")
		mes, err := tf.ReadPkg()
		if err != nil {
			fmt.Println("服务器出错，无法接收信息", err)
			return
		}
		fmt.Println("mes = ", mes)
	}
}
