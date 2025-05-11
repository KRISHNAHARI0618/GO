package main

import "fmt"

type Person struct {
	Name string
	age  int
}

func (p Person) Speak() {
	fmt.Println("My name is ", p.Name)
	fmt.Println("My age is  ", p.age)
}

func main() {

	p1 := Person{
		Name: "Peddireddy",
		age:  24,
	}
	p2 := Person{
		Name: "Hari Vardhan",
		age:  25,
	}

	p1.Speak()
	p2.Speak()

}
