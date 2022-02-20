package main

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

type User struct {
	UserId   int    `json:"userId"`
	UserPwd  string `json:"userPwd"`
	UserName string `json:"userName"`
	UserInfo Info
}

type Info struct {
	Mes string `json:"mes"`
	Err error  `json:"err"`
}

func main() {
	conn, err := redis.Dial("tcp", "112.124.37.136:6379")
	if err != nil {
		fmt.Println("", err)
		return
	}

	_, err = conn.Do("hset", "users", 100, "{\"userId\":100,\"userPwd\":\"123456\",\"userName\":\"jack\"}")
	if err != nil {
		fmt.Println("", err)
		return
	}

	res, err := redis.String(conn.Do("hget", "users", 100))
	if err != nil {
		fmt.Println("", err)
		return
	}
	fmt.Println(res)

	// var user User
	// json.Unmarshal([]byte(res), &user)
	// fmt.Println(user)
	// info := Info{
	// 	Mes: "123",
	// 	Err: errors.New("24"),
	// }
	// user := User{
	// 	UserId:   1,
	// 	UserPwd:  "123",
	// 	UserName: "jack",
	// 	UserInfo: info,
	// }
	// res, err := json.Marshal(user)
	// if err != nil {
	// 	fmt.Println("", err)
	// 	return
	// }
	// fmt.Println(string(res))

}
