package System

import (
	"fmt"
	"io"
	"net"
	"sync"
	"time"
)

type Server struct {
	IP   string
	Port int
	// 在线用户的列表
	OnlineMap map[string]*User
	MpLock    sync.RWMutex

	//消息广播channel
	Message chan string
}

func NewServer(ip string, port int) *Server {

	return &Server{
		IP:        ip,
		Port:      port,
		OnlineMap: make(map[string]*User),
		Message:   make(chan string),
	}
}

// Handle 处理连接的业务
func (server *Server) Handle(conn net.Conn) {
	fmt.Println(conn.RemoteAddr().String() + "连接成功")
	// 用户上线 将用户加入 onlinemap中
	user := NewUser(conn, server)
	user.Online()
	// 监听用户是否活跃的channel
	isLive := make(chan bool)
	go func() {
		byteValus := make([]byte, 4096)
		for {
			n, err := conn.Read(byteValus)
			if err != nil && err != io.EOF {
				fmt.Println("conn read err:", err)
			}
			if n == 0 {
				user.Offline()
				return
			}
			// 获取用户的消息('\n')
			msg := string(byteValus[:n])
			user.DoMessage(msg)
			//用户的任意消息,代表当前用户是活跃的
			isLive <- true
		}
	}()

	//当前handler阻塞 使 协程不能死亡
	for {
		select {
		case <-isLive:
			//当前用户是活跃的，应该重置定时器
			// 不做任何事，只是为了激活定时器
		case <-time.After(time.Hour):
			//超时 将当前的user强制关闭
			user.SendMsg("你被踢了")
			//关闭当前用户的管道
			close(user.C)
			// 关闭当前用的连接
			conn.Close()
			//退出当前的handler
			return

		}
	}

}
func (server *Server) ListenMesage() {
	for {
		msg := <-server.Message
		server.MpLock.Lock()
		for _, cli := range server.OnlineMap {
			cli.C <- msg
		}
		server.MpLock.Unlock()
	}
}

// BroadCast 消息广播
func (server *Server) BroadCast(user *User, msg string) {
	sendMsg := "[" + user.Addr + "]" + user.Name + ":" + msg
	fmt.Println(sendMsg)
	// 想Message 中发送消息  所以应该有一个监听MessAge的方法
	server.Message <- sendMsg
}

// Start 服务端连接
func (server Server) Start() {
	// Socket Listen
	liesten, err := net.Listen("tcp", fmt.Sprintf("%s:%d", server.IP, server.Port))
	if err != nil {
		fmt.Println("net Listen err", err)
		return
	}
	defer liesten.Close()
	fmt.Println("监听开始")
	// 启动监听ListenMessage
	go server.ListenMesage()
	for {
		//Accept
		conn, err := liesten.Accept()
		if err != nil {
			fmt.Println("Listen accept err", err)
			continue
		}
		defer conn.Close()
		// do handler
		go server.Handle(conn)
	}

	//close Listen socket
}
