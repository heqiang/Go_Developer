package System

import (
	"fmt"
	"net"
	"strings"
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
	user.server.BroadCast(user, "上线了\n")
}

//用户下线
func (user *User) Offline() {
	user.server.MpLock.Lock()
	delete(user.server.OnlineMap, user.Name)
	user.server.MpLock.Unlock()
	user.server.BroadCast(user, "下线了\n")
}

// 向当前cleient 对应的用户发送消息
func (user *User) SendMsg(msg string) {
	user.conn.Write([]byte(msg))
}

// 用户处理消息的业务
func (user *User) DoMessage(msg string) {
	trimSpaceMsg := strings.TrimSpace(msg)
	if trimSpaceMsg == "who" {
		//用户在线查询
		user.server.MpLock.Lock()
		for _, user1 := range user.server.OnlineMap {
			// 去除当前客户端的在线信息
			fmt.Println("用户姓名:", user1.Name, "user姓名：", user.Name)
			if strings.TrimSpace(user.Name) != strings.TrimSpace(user1.Name) {
				onlineMsg := "[" + user1.Addr + "]" + strings.TrimSpace(user1.Name) + ":" + "在线\n"
				user.SendMsg(onlineMsg)
			}
		}
		user.server.MpLock.Unlock()
		//用户名更改
	} else if strings.Contains(trimSpaceMsg, "rename") {
		name := strings.Contains(trimSpaceMsg, "#")
		if name {
			newName := strings.Split(trimSpaceMsg, "#")[1]
			if len(strings.TrimSpace(newName)) > 0 {
				//判断用户名是否被占用
				if _, ok := user.server.OnlineMap[newName]; ok {
					msg := fmt.Sprintf("%s被占用", newName)
					user.SendMsg(msg)
				} else {
					user.server.MpLock.Lock()
					delete(user.server.OnlineMap, user.Name)
					fmt.Println()
					user.server.OnlineMap[newName] = user
					user.server.MpLock.Unlock()
					user.Name = newName
					user.SendMsg("修改成功\n")
				}
			} else {
				user.SendMsg("用户名不能为空,请从新输入\n")
			}
		} else {
			msg := "输入的格式有误 正确格式 rename#newName\n"
			user.SendMsg(msg)
		}

	} else {
		// 用户群聊广播
		user.server.BroadCast(user, msg)
	}

}
