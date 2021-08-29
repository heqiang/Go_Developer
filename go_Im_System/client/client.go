package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
)

type Client struct {
	ServerIp   string
	ServerPort int
	Name       string
	conn       net.Conn
	Flag       int
}

func NewClient(serverip string, serverport int) *Client {
	client := &Client{
		ServerIp:   serverip,
		ServerPort: serverport,
		Flag:       999,
	}
	address := fmt.Sprintf("%s:%d", client.ServerIp, client.ServerPort)
	conn, err := net.Dial("tcp", address)
	if err != nil {
		fmt.Println("net dial err:", err)
	}
	client.conn = conn

	return client
}

var serverip string
var serverport int

// ./client -ip 127.0.0.1 -port 8888

func init() {
	flag.StringVar(&serverip, "ip", "127.0.0.1", "设置服务器地址(默认127.0.0.1)")
	flag.IntVar(&serverport, "port", 8888, "设置服务器端口(默认8888)")
}

func (client *Client) menu() bool {
	var flag int

	fmt.Println("1 公聊模式,2 私聊模式，3 更新用户名,0 退出")
	fmt.Scanln(&flag)

	if flag >= 0 && flag <= 3 {
		client.Flag = flag
		return true
	} else {
		fmt.Println("》》》请输入合法范围内的数字《《《")
		return false
	}
}
func (client *Client) RunClient() {
	for client.Flag != 0 {
		for client.menu() != true {

		}
		switch client.Flag {
		case 1:
			//公聊模式
			fmt.Println("公聊模式选择")
			client.PublishChat()
			break
		case 2:
			//私聊模式
			fmt.Println("私聊模式选择")
			client.PrivatChat()
			break
		case 3:
			// 更新用户名
			client.UpdateName()
		}
	}
}
func (client *Client) UpdateName() bool {
	fmt.Println(">>>>>>>>请输入用户名")
	fmt.Scanln(&client.Name)
	sendMsg := "rename#" + client.Name + "\n"
	_, err := client.conn.Write([]byte(sendMsg))
	if err != nil {
		fmt.Scanln("修改失败")
		return false
	}
	return true
}

//处理server回应的消息，直接显示到标准输出即可
func (client *Client) DealResponse() {
	//一旦有client.conn 有数据 就直接copy到stdout的标准输出上，永久阻塞监听
	io.Copy(os.Stdout, client.conn)
}
func (client *Client) PublishChat() {
	var chatMsg string

	fmt.Println(">>>> 请输入聊天内容，exit退出")
	fmt.Scanln(&chatMsg)

	for chatMsg != "exit" {
		// 发送消息给服务器
		if len(strings.TrimSpace(chatMsg)) != 0 {
			_, err := client.conn.Write([]byte(chatMsg + "\n"))
			if err != nil {
				fmt.Println("client write err:", err)
				break
			}
		}
		chatMsg = ""
		fmt.Println(">>>> 请输入聊天内容，exit退出")
		fmt.Scanln(&chatMsg)
	}

}

//查询在线用户
func (client *Client) SelectUser() {
	sendMsg := "who\n"
	_, err := client.conn.Write([]byte(sendMsg))
	if err != nil {
		fmt.Println("conn write err:", err)
		return
	}

}

//私聊模式的实现
func (client *Client) PrivatChat() {
	var remoteName string
	var chatMsg string

	client.SelectUser()
	fmt.Println(">>>> 请输入聊天内容，exit退出")
	fmt.Scanln(&remoteName)

	for remoteName != "exit" {
		fmt.Println(">>>> 请输入聊天内容，exit退出")
		fmt.Scanln(&chatMsg)
		for chatMsg != "exit" {
			if len(strings.TrimSpace(chatMsg)) != 0 {
				sendMsg := "to#" + remoteName + "#" + chatMsg + "\n"
				_, err := client.conn.Write([]byte(sendMsg + "\n"))
				if err != nil {
					fmt.Println("client write err:", err)
					break
				}
			}
			chatMsg = ""
			fmt.Println(">>>> 请输入聊天内容，exit退出")
			fmt.Scanln(&chatMsg)
		}
	}
}

func main() {
	//命令行解析
	flag.Parse()
	client := NewClient("127.0.0.1", 8888)
	//client := NewClient(serverip,serverport)
	if client == nil {
		fmt.Println("服务器连接失败")
	}
	fmt.Println("服务器连接成功")
	//单独开启一个协程处理server的response
	go client.DealResponse()

	client.RunClient()
}
