package model

import (
	"chatroom/common/message"
	"encoding/json"
	"fmt"

	"github.com/garyburd/redigo/redis"
)

//在服务器启动后，就初始化一个UserDao实例
//作为全局变量，在需要数据库操作时，直接使用
var (
	MyUserDao *UserDao
)

//定义一个UserDao结构体，完成对User结构体的数据库操作
type UserDao struct {
	pool *redis.Pool
}

//使用工厂模式，创建一个UserDao实例
func NewUserDao(pool *redis.Pool) (userDao *UserDao) {
	userDao = &UserDao{
		pool: pool,
	}
	return
}

//1. 根据一个用户ID，返回一个用户实例和err
func (this *UserDao) GetUserById(conn redis.Conn, id int) (user *User, err error) {
	fmt.Printf("用户ID=%d", id)
	res, err := redis.String(conn.Do("hget", "users", id))
	fmt.Println("得到res", res)
	if err != nil {
		if err == redis.ErrNil {
			err = ERROR_USER_NOTEXITS
		}
		return
	}
	//反序列化得到User实例
	err = json.Unmarshal([]byte(res), &user)
	if err != nil {
		fmt.Println("反序列化得到User出错", err)
		return
	}
	fmt.Println("得到用户", user)
	return
}

func (this *UserDao) Login(userId int, userPwd string) (user *User, err error) {
	//先从UserDao连接池中取出一个连接
	conn := this.pool.Get()
	defer conn.Close()
	user, err = this.GetUserById(conn, userId)
	if err != nil {
		return
	}
	//用户存在，接下来检验密码
	if user.UserPwd != userPwd {
		err = ERROR_USER_PWD
		return
	}
	return
}

func (this *UserDao) Register(user *message.User) (err error) {
	//先从UserDao连接池中取出一个连接
	conn := this.pool.Get()
	defer conn.Close()
	_, err = this.GetUserById(conn, user.UserId)
	//用户已存在的错误
	if err == nil {
		err = ERROR_USER_EXITS
		return
	}
	//用户不存在，可以创建
	data, err := json.Marshal(user)
	if err != nil {
		fmt.Println("序列化用户出错", err)
		return
	}

	//入库
	_, err = conn.Do("hset", "users", user.UserId, string(data))

	if err != nil {
		fmt.Println("注册用户入库错误=", err)
		return
	}
	return
}
