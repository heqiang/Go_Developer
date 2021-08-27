package main

import "fmt"

type  People struct {
	Name string
	Age int
}

func  (p People) Show() {
	fmt.Println("Name:",p.Name)
	fmt.Println("Age:",p.Age)
}

func  (p People)GetName() string {
		return  p.Name
}

func  (p *People)SetName(name string)  {
	  p.Name = name
}

func  main()  {
	 p:= People{
	 	Name: "hq",
	 	Age: 16,
	 }
	 p.Show()
	 name:=p.GetName()
	 fmt.Println(name)
	 fmt.Println("************")
	 p.SetName("jjjj")
	 name1:=p.GetName()
	 fmt.Println(name1)
	 p.Show()
}
