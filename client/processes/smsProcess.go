package processes

import (
	"chatroom/common/message"
	"chatroom/server/utils"
	"encoding/json"
	"fmt"
)

type SmsProcess struct {
}

//发送群聊消息
func (this *SmsProcess) SendGroupMes(content string) (err error) {
	//创建一个Mes
	var mes message.Message
	mes.Type = message.SmsMesType
	//创建一个SmsMes
	var smsMes message.SmsMes
	smsMes.Content = content
	smsMes.UserId = CurUser.UserId
	smsMes.UserStatus = CurUser.UserStatus

	//序列化smsMes
	data, err := json.Marshal(smsMes)
	if err != nil {
		fmt.Println("序列化smsMes出错", err)
		return
	}

	mes.Date = string(data)

	//序列化Mes
	data, err = json.Marshal(smsMes)
	if err != nil {
		fmt.Println("序列化smsMes第二次出错", err)
		return
	}

	//发送序列化后的mes
	tf := &utils.Transfer{
		Conn: CurUser.Conn,
	}

	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("发送smsMes失败", err)
		return
	}
	return
}
