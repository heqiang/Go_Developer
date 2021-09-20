package main

import (
	"Go_Developer/go_redis/initConn"
	"Go_Developer/go_redis/redis_watch"
	"fmt"
)

func main() {
	if err := initConn.InitClient(); err != nil {
		fmt.Println("init client failed.err:%v\n", err)
		return
	}
	defer initConn.Rdb.Close()

	//reidis_pipeline.Ceshi(initConn.Rdb)
	//redis_txtpipeline.TxPieline(initConn.Rdb)
	redis_watch.Watch1(initConn.Rdb)
}
