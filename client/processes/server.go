package processes

import (
	"chatroom/client/utils"
	"chatroom/common/message"
	"encoding/json"
	"fmt"
	"net"
	"os"
)

//显示登录成功后的界面
func ShowMenu() {
	fmt.Println("----------恭喜XXX登录成功----------")
	fmt.Println("----------1. 在线用户----------")
	fmt.Println("----------2. 发送消息----------")
	fmt.Println("----------3. 消息列表----------")
	fmt.Println("----------4. 退出系统----------")
	fmt.Print("请选择(1-4): ")
	var key int
	var content string

	//因为总会用到SmsProcess实例，所以定义外switch外部
	//但实际上循环在这个函数外，所以还是每次都创建了
	smsProcess := &SmsProcess{}
	fmt.Scanf("%d\n", &key)
	switch key {
	case 1:
		outputOnlineUsers()
	case 2:
		fmt.Println("你想群发什么信息: ")
		fmt.Scanf("%s\n", content)
		smsProcess.SendGroupMes(content)
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
		switch mes.Type {
		case message.NotifyUserStatusMesType:
			//有人上线了
			//1. 取出notifyUserStatusMes
			var notifyUserStatusMes message.NotifyUserStatusMes
			json.Unmarshal([]byte(mes.Date), &notifyUserStatusMes)
			//2. 把新上线用户加入到客户端维护的map中
			updateUserStatus(&notifyUserStatusMes)
		case message.SmsMesType:
			outputGroupMes(&mes)
		default:
			fmt.Println("服务器发送了未知的消息类型")
		}
	}
}
