package main

import (
	"Go_Developer/go_consul/pb"
	"context"
	"fmt"
	"google.golang.org/grpc"
)

func main() {
	//1 服务连接
	grpcConn, _ := grpc.Dial("127.0.0.1:8889", grpc.WithInsecure())
	//2 初始化grpc客户端
	grpcClient := pb.NewHelloClient(grpcConn)
	var p pb.Person
	p.Name = "hq"
	p.Age = 12
	//远程函数的调用
	p1, err := grpcClient.SayHello(context.TODO(), &p)
	fmt.Println(p1.Age, p1.Name, err)
}
