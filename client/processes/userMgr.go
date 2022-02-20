package processes

import (
	"chatroom/client/model"
	"chatroom/common/message"
	"fmt"
)

//客户端维护的在线用户列表
//初始化在用户刚登录时，从服务器收到目前在线用户的切片
var onlineUsers map[int]*message.User = make(map[int]*message.User)

var CurUser model.CurUser //在用户登录后完成初始化

func outputOnlineUsers() {
	//遍历onlineUsers
	fmt.Println("当前在线用户列表:")
	for id, _ := range onlineUsers {
		fmt.Printf("用户id:\t%d", id)
	}
}

//
func updateUserStatus(notifyUserStatusMes *message.NotifyUserStatusMes) {

	//
	user, ok := onlineUsers[notifyUserStatusMes.UserId]
	if !ok {
		// 原来没有
		user = &message.User{
			UserId: notifyUserStatusMes.UserId,
		}
	}
	user.UserStatus = notifyUserStatusMes.Status

	onlineUsers[notifyUserStatusMes.UserId] = user

	outputOnlineUsers()
}
