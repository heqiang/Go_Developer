package main

import "fmt"

type  Animal interface {
	Sleep()
	GetColor()  string
	GetType()  string
}

type  Cat struct {
	Color string
	Type  string
}

func  (c Cat) Sleep() {
	fmt.Println("cat is sleep ")
}
func  (c Cat) GetColor()  string{
	return  c.Color
}
func  (c Cat) GetType()  string {
	return c.Type
}
// 另一个具体的类
type  Dog struct {
	Color string
	Type  string
}
func  (c Dog) Sleep() {
	fmt.Println("Dog is sleep ")
}
func  (c Dog) GetColor()  string{
	return  c.Color
}
func  (c Dog) GetType()  string {
	return c.Type
}
func main()  {
	var animal Animal
	animal = &Cat{
		Color: "黄色",
		Type: "拉布拉多",
	}
	animal1 := &Dog{
		Color: "白色",
		Type: "牧羊犬",
	}
	fmt.Println(animal.GetColor())
	fmt.Println(animal1.GetColor())
}


