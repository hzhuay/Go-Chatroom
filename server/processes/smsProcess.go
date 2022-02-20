package processes

import (
	"chatroom/common/message"
	"chatroom/server/utils"
	"encoding/json"
	"fmt"
	"net"
)

type SmsProcess struct{}

func (this *SmsProcess) SendGroupMes(mes message.Message) {
	//遍历服务器端的onlineUsers，群发信息

	var smsMes message.SmsMes
	err := json.Unmarshal([]byte(mes.Date), &smsMes)
	if err != nil {
		fmt.Println("SendGroupMes=", err)
		return
	}

	data, err := json.Marshal(mes)
	if err != nil {
		fmt.Println("SendGroupMes中序列化", err)
		return
	}

	for id, up := range userMgr.onlineUsers {
		//过滤自己
		if id == smsMes.UserId {
			continue
		}
		this.SendMesToEachOnlineUser(data, up.Conn)
	}
}

func (this *SmsProcess) SendMesToEachOnlineUser(data []byte, conn net.Conn) {
	tf := &utils.Transfer{
		Conn: conn,
	}
	err := tf.WritePkg(data)
	if err != nil {
		fmt.Println("SendMesToEachOnlineUser", err)
		return
	}
}
