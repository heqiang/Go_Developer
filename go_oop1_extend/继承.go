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

func  (s People)Eat(){
	fmt.Println("son:"+s.Name)
}

type  SurperPeople struct {
	People
	Level  int
}

func  (s SurperPeople)Eat(){
	fmt.Println("surper:"+s.Name)
}

func  (p SurperPeople) Show() {
	fmt.Println("Name:",p.Name)
	fmt.Println("Age:",p.Age)
}


func  main()  {
	surper:=SurperPeople{
		People{
			Name: "HQ",
			Age: 16,
		},12,
	}
	var  s  SurperPeople
	s.Name = "sss"
	s.Age = 18
	s.Level = 33
	surper.Eat()
	s.Show()

}
