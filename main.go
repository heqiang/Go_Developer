package main

import (
	"fmt"
	"os"
	"runtime"
	"time"
)

func main() {
	//go CeshiTime()
	//
	//time.Sleep(30*time.Second)
	//str := "to#姓名# "
	//res := strings.Split(strings.TrimSpace(str), "#")
	//fmt.Println(res)
	//fmt.Println(len(res))

	_, fileStr, _, _ := runtime.Caller(1)
	fmt.Println(fileStr)
	dir, _ := os.Getwd()
	fmt.Println(dir)
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
