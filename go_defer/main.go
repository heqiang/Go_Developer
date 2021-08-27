package main

import "fmt"

func  deferFunc()  {
	fmt.Println("defer 执行了")
}

func  returnFunc() int {
	fmt.Println("return 执行了")
	return 0
}

func  ceshi() int {

	defer  deferFunc()

	return returnFunc()
}


func  main()  {
	ceshi()
}
