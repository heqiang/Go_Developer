package main

import (
	"Go_Developer/go_consul/pb"
	"context"
	"fmt"
	"github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
	"net"
)

//定义类
type Child struct {
}

//绑定类方法，实现接口
func (child *Child) SayHello(ctx context.Context, person *pb.Person) (*pb.Person, error) {
	person.Name = "HELLO" + person.Name
	return person, nil
}

func main() {
	// 将grpc服务注册到consul上
	consulConfig := api.DefaultConfig()
	//创建consul对象
	consulClient, err := api.NewClient(consulConfig)
	if err != nil {
		fmt.Println("api NewClient err:", err)
		return
	}
	//告诉consul 即将注册的服务的配置信息
	registerService := &api.AgentServiceRegistration{
		ID: "bj38",
		//创建服务的别名
		Tags:    []string{"grpc", "consul"},
		Name:    "grpc And consul",
		Address: "127.0.0.1",
		Port:    8889,
		Check: &api.AgentServiceCheck{
			CheckID:  "consul grpc test",
			TCP:      "127.0.0.1:8889",
			Timeout:  "1s",
			Interval: "5s",
		},
	}
	//注册grpc服务到consul上
	consulClient.Agent().ServiceRegister(registerService)

	//////////////////////
	//初始化 grpc服务
	grpcServer := grpc.NewServer()

	//注册服务
	pb.RegisterHelloServer(grpcServer, new(Child))

	listener, err := net.Listen("tcp", "127.0.0.1:8889")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer listener.Close()
	fmt.Println(">>>>> 服务启动成功")
	//启动服务
	grpcServer.Serve(listener)

}
