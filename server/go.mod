module chatroom/server

go 1.17

require (
	chatroom/common v0.0.0
	chatroom/server/model v0.0.0-00010101000000-000000000000
	github.com/garyburd/redigo v1.6.2
)

replace chatroom/common => ../common //本地包相对路径或绝对路径

replace chatroom/server/model => ./model //本地包相对路径或绝对路径
