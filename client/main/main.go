package main

import (
	"chatroom/client/processes"
	"fmt"
)

var userID int
var userPwd string
var userName string

func main() {

	//接收用户选择
	var key int
	//判断是否继续显示菜单
	// var loop bool = true
	for true {
		fmt.Println("----------欢迎登录多人聊天系统----------")
		fmt.Println("\t\t 1 登录聊天室")
		fmt.Println("\t\t 2 注册用户")
		fmt.Println("\t\t 3 退出系统")
		fmt.Println("\t\t 请选择(1-3)")

		fmt.Scanf("%d\n", &key)

		switch key {
		case 1:
			fmt.Println("登录聊天室")
			fmt.Print("请输入用户的ID: ")
			fmt.Scanf("%d\n", &userID)
			fmt.Print("请输入用户的密码: ")
			fmt.Scanf("%s\n", &userPwd)

			// 创建一个UserProcess实例
			up := &processes.UserProcess{}
			up.Login(userID, userPwd)
			// loop = false
		case 2:
			fmt.Println("注册用户")
			fmt.Print("请输入用户的ID: ")
			fmt.Scanf("%d\n", &userID)
			fmt.Print("请输入用户的密码: ")
			fmt.Scanf("%s\n", &userPwd)
			fmt.Print("请输入用户的昵称: ")
			fmt.Scanf("%s\n", &userName)
			// loop = false
			//创建一个UserProcess实例，完成注册
			up := &processes.UserProcess{}
			up.Register(userID, userPwd, userName)
		case 3:
			fmt.Println("退出系统")
			return
		default:
			fmt.Println("您的输入有误，请重新输入")
		}
	}

}
