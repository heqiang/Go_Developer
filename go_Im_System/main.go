package main

import "Go_Developer/go_Im_System/ImServer"

func  main()  {
	server := ImServer.NewServer("127.0.0.1",8888)
	server.Start()
}