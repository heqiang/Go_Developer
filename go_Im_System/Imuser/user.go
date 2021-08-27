package Imuser

import (
	"net"
)

type User struct {
	Name  string
	Addr  string
	C chan string
	conn net.Conn
}

// 创建一个用户的api
func  NewUser(conn net.Conn) *User  {
	userAddr := conn.RemoteAddr().String()
	user:=&User{
		Name: userAddr,
		Addr: userAddr,
		C: make(chan  string),
		conn:  conn,
	}
	go user.listenMessage()
	return  user
}
// 监听当前 user channel 的方法 ，一旦有消息 就直接发给客户端
func  (user  User)listenMessage()  {
	 for{
	 	msg:= <-user.C
	 	user.conn.Write([]byte(msg+"\n"))
	 }
}
