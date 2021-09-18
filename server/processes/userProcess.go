package processes

import (
	"chatroom/common/message"
	"chatroom/server/model"
	"chatroom/server/utils"
	"encoding/json"
	"fmt"
	"net"
)

type UserProcess struct {
	Conn net.Conn
}

func (this *UserProcess) ServerProcessLogin(mes *message.Message) (err error) {
	//从mes中取出data，并且反序列化
	var loginMes message.LoginMes
	err = json.Unmarshal([]byte(mes.Date), &loginMes)
	if err != nil {
		fmt.Println("json.Unmarshal fail err = ", err)

	}

	//定义要回应的信息
	var resMes message.Message
	resMes.Type = message.LoginResMesType
	var loginResMes message.LoginResMes

	//去Redis验证用户
	user, err := model.MyUserDao.Login(loginMes.UserId, loginMes.UserPwd)
	if err != nil {
		if err == model.ERROR_USER_NOTEXITS {
			loginResMes.Code = 500

		} else if err == model.ERROR_USER_PWD {
			loginResMes.Code = 403
		} else {
			//未知错误
			loginResMes.Code = 505
		}
		loginResMes.Error = err.Error()
		//先让测试成功
	} else {
		loginResMes.Code = 200
		fmt.Println(user, "登录成功")
	}

	//序列化
	data, err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Println("json.Marshal fail ", err)
	}
	resMes.Date = string(data)

	//对resMes序列化
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.Marshal fail ", err)
	}

	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	err = tf.WritePkg(data)
	return
}

func (this *UserProcess) ServerProcessRegister(mes *message.Message) (err error) {
	return
}
