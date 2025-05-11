package main

import "fmt"

type person struct {
	first string
}
type secretagent struct {
	person
	ltk bool
}

func (p person) speak() {
	fmt.Println("I am Calling Man", p.first)
}

func (sa secretagent) speak() {
	fmt.Println("My Name is :", sa.first)
}

type human interface {
	speak()
}

func saysomething(h human) {
	h.speak()
}

func main() {

	sa1 := secretagent{
		person: person{
			first: "peddireddy",
		},
		ltk: true,
	}
	saysomething(sa1)

}
