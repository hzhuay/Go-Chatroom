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
	Conn   net.Conn
	UserId int
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

		//用户登陆成功后，把id加入UserProcess中，再放入UserMgr中
		this.UserId = loginMes.UserId
		userMgr.AddOnlineUser(this)
		this.NotifyOtherOnlineUsers()
		//将在线用户的id加入loginResMes.Usersid
		for id, _ := range userMgr.onlineUsers {
			loginResMes.UsersId = append(loginResMes.UsersId, id)
		}

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
	//从mes中取出data，并且反序列化
	fmt.Println("进入ServerProcessRegister")

	var registerMes message.RegisterMes
	err = json.Unmarshal([]byte(mes.Date), &registerMes)
	if err != nil {
		fmt.Println("json.Unmarshal fail err = ", err)

	}

	//定义要回应的信息
	var resMes message.Message
	resMes.Type = message.RegisterResMesType
	var registerResMes message.RegisterResMes

	//去Redis添加用户
	err = model.MyUserDao.Register(&registerMes.User)
	if err != nil {
		if err == model.ERROR_USER_EXITS {
			registerResMes.Code = 505
			registerResMes.Error = model.ERROR_USER_EXITS.Error()
		} else {
			registerResMes.Code = 506
			registerResMes.Error = "注册发生未知错误"
		}
	} else {
		registerResMes.Code = 200

	}

	//序列化
	data, err := json.Marshal(registerResMes)
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

//通知用户上线
//按理说userId不需要作为参数，在this里
func (this *UserProcess) NotifyOtherOnlineUsers() {
	//遍历onlineUsers
	for id, up := range userMgr.onlineUsers {
		if id == this.UserId {
			continue
		}
		//调用其他用户的UserProcess，把新上线用户的ID传过去
		up.NotifyMeOnline(this.UserId)
	}
}

func (this *UserProcess) NotifyMeOnline(userId int) {
	//组装NotifyUserStatusMes
	var mes message.Message
	mes.Type = message.NotifyUserStatusMesType

	var notifyUserStatusMes message.NotifyUserStatusMes
	notifyUserStatusMes.UserId = userId
	notifyUserStatusMes.Status = message.UserOnline

	//将notifyUserStatusMes序列化
	data, err := json.Marshal(notifyUserStatusMes)
	if err != nil {
		fmt.Println("序列化notifyUserStatusMes出错=", err)
		return
	}
	mes.Date = string(data)

	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("序列化notifyUserStatusMes第二次出错=", err)
		return
	}
	fmt.Println("给其他用户发送")
	//发送
	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("发送用户状态改变包", err)
		return
	}

}
