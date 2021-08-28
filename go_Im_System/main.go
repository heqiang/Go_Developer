package main

import (
	"Go_Developer/go_Im_System/System"
)

func main() {
	server := System.NewServer("127.0.0.1", 8888)
	server.Start()
}
