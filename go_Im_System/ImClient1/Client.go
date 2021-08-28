package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("net 连接 失败 err:", err)
		return
	}
	defer conn.Close()
	bytevalue := make([]byte, 32)
	_, err = conn.Read(bytevalue)
	if err != nil {
		fmt.Println("client read err:", err)
		return
	}
	fmt.Println(string(bytevalue))

}
