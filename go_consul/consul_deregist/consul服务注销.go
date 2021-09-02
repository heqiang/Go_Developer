package main

import (
	"fmt"
	"github.com/hashicorp/consul/api"
)

func main() {
	consulConfig := api.DefaultConfig()

	consulClinet, err := api.NewClient(consulConfig)
	if err != nil {
		fmt.Println("err,", err)
	}
	consulClinet.Agent().ServiceDeregister("bj38")
}
