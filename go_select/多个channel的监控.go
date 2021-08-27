package main

import "fmt"

func  Fib(c,quit chan int )  {
	x,y:=1,1

	for {
		select {
		case c<-x:
			x,y = y,x+y
		case <-quit:
			fmt.Println("结束")
			return

		}
	}
}

func  main()  {
	c:= make(chan int)
	quit:= make(chan int)
	go func() {
		for i:=0;i<6;i++{
			fmt.Println(<-c)
		}
		quit<-0
	}()

	Fib(c,quit)
}

