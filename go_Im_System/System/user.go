package System

import (
	"net"
)

type User struct {
	Name   string
	Addr   string
	C      chan string
	conn   net.Conn
	server *Server
}

// 创建一个用户的api
func NewUser(conn net.Conn, server *Server) *User {
	userAddr := conn.RemoteAddr().String()
	user := &User{
		Name:   userAddr,
		Addr:   userAddr,
		C:      make(chan string),
		conn:   conn,
		server: server,
	}
	go user.listenMessage()
	return user
}

// 监听当前 user channel 的方法 ，一旦有消息 就直接发给客户端
func (user User) listenMessage() {
	for {
		msg := <-user.C
		user.conn.Write([]byte(msg + "\n"))
	}
}

// 用户上线
func (user *User) Online() {

	user.server.MpLock.Lock()
	user.server.OnlineMap[user.Name] = user
	user.server.MpLock.Unlock()
	user.server.BroadCast(user, "上线了")
}

//用户下线
func (user *User) Offline() {
	user.server.MpLock.Lock()
	delete(user.server.OnlineMap, user.Name)
	user.server.MpLock.Unlock()
	user.server.BroadCast(user, "下线了")
}

// 用户处理消息的业务
func (user *User) DoMessage(msg string) {
	user.server.BroadCast(user, msg)

}
