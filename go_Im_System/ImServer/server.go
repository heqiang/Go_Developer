package ImServer

import (
	"Go_Developer/go_Im_System/Imuser"
	"fmt"
	"net"
	"sync"
)

type  Server struct {
	IP string
	Port  int
	// 在线用户的列表
	OnlineMap  map[string]*Imuser.User
	MpLock  sync.RWMutex

	//消息广播channel
	Message  chan string
}

func NewServer(ip string,port int) *Server {

	return  &Server{
		 IP: ip,
		 Port: port,
		 OnlineMap: make(map[string]*Imuser.User),
		 Message : make(chan string),
	}
}

// Handle 处理连接的业务
func  (server *Server) Handle(conn net.Conn) {
	fmt.Println("连接成功")
	// 用户上线 将用户加入 onlinemap中
	user:=Imuser.NewUser(conn)
	server.MpLock.Lock()
	server.OnlineMap[user.Name] = user
	server.MpLock.Unlock()

	server.BroadCast(user,"上线了")

	select {

	}

}
func  (server *Server)ListenMesage()  {
	for {
		msg:=<-server.Message
		server.MpLock.Lock()
		for _,cli:=range  server.OnlineMap{
			cli.C <- msg
		}

	}
}


// BroadCast 消息广播
func  (server *Server)BroadCast(user *Imuser.User,msg string)  {
		sendMsg := "[" + user.Addr+"]"+user.Name+":" +msg
		// 想Message 中发送消息  所以应该有一个监听MessAge的方法
		server.Message <- sendMsg
}

// Start 服务端连接
func  (server Server)Start()  {
	// Socket Listen
	liesten,err:=net.Listen("tcp",fmt.Sprintf("%s:%d",server.IP,server.Port))
	if err!=nil{
		fmt.Println("net Listen err",err)
		return
	}
	defer  liesten.Close()
	fmt.Println("监听开始")
	// 启动监听ListenMessage
	go server.ListenMesage()
	for {
		//Accept
		conn,err:= liesten.Accept()
		if err!=nil{
			fmt.Println("Listen accept err",err)
			continue
		}
		defer conn.Close()
		// do handler
		fmt.Println("")
		go server.Handle(conn)
	}


	//close Listen socket
}