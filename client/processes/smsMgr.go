package processes

import (
	"chatroom/common/message"
	"encoding/json"
	"fmt"
)

func outputGroupMes(mes *message.Message) {
	//反序列化mes.Data

	var smsMes message.SmsMes
	err := json.Unmarshal([]byte(mes.Date), &smsMes)
	if err != nil {
		fmt.Println("outputGroupMes反序列化", err)
		return
	}

	//显示
	fmt.Printf("用户id:\t%d 对大家说:\t%s\n", smsMes.UserId, smsMes.Content)

}
