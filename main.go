package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	//go CeshiTime()
	//
	//time.Sleep(30*time.Second)
	str := "rename#hh"
	res := strings.Split(str, "#")[1]
	fmt.Println(res)
}

func CeshiTime() {
	isLive := make(chan bool)

	go Msg(isLive)
	for {
		select {
		case <-isLive:
			fmt.Println("当前用户活跃中")
		case <-time.After(5 * time.Second):
			fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
			fmt.Println("时间到了")
			return
		}
	}

}

func Msg(live chan bool) {
	for i := 0; i < 5; i++ {
		fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
		time.Sleep(3 * time.Second)
		live <- true
	}
}
