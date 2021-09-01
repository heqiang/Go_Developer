package main

import (
	"Go_Developer/go_consul/pb"
	"context"
	"fmt"
	"github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
	"strconv"
)

func main() {
	//初始化consul配置 客户端服务器需要一致
	consulConfig := api.DefaultConfig()

	//获取consul的操作对象 -- (可以重新指定consul 属性IP/Port 也可以使用默认)
	registerClient, err := api.NewClient(consulConfig)
	//获取地址
	if err != nil {
		fmt.Println("client NewClient err:", err)
		return
	}
	//从caonsul上获取健康的服务,
	// params
	// service 注册的name
	//tag 注册的Tag 有多个的时候随意选择一个
	//passingOnly 是否通过健康检查
	//QueryOptions 通常传nill
	// return []*ServiceEntry 注册的所有的服务集群切片 QueryMeta额外查询返回值 一般不需要 error 错误
	serviceEntry, _, _ := registerClient.Health().Service("grpc And consul", "consul", true, nil)
	address := serviceEntry[0].Service.Address
	port := serviceEntry[0].Service.Port
	target := address + ":" + strconv.Itoa(port)
	grpcConn, _ := grpc.Dial(target, grpc.WithInsecure())

	//1 服务连接
	//grpcConn, _ := grpc.Dial("127.0.0.1:8889", grpc.WithInsecure())
	////2 初始化grpc客户端
	grpcClient := pb.NewHelloClient(grpcConn)
	var p pb.Person
	p.Name = "hq"
	p.Age = 12
	//远程函数的调用
	p1, err := grpcClient.SayHello(context.TODO(), &p)
	if err != nil {
		fmt.Println("grpcClient.SayHello err:", err)
		return
	}
	fmt.Println(p1.Age, p1.Name, err)
}
