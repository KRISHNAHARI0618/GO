package main

import "fmt"

type person1 struct {
	name   string
	age    int
	marks  float64
	isPass bool
}

type person2 struct {
	name   string
	age    int
	marks  float64
	isPass bool
}

func (p1 person1) student() {
	fmt.Println("I am Calling Eddulamma")

	fmt.Printf("My name is : %+v , My Age is : %v , My Marks is: %v ,I am passed: %v ", p1.name, p1.age, p1.marks, p1.isPass)

	fmt.Println()
}

func (p2 person2) student() {
	fmt.Println("I am calling Muddulamma ")

	fmt.Printf("My name is : %+v , My Age is : %v , My Marks is: %v ,I am passed: %v ", p2.name, p2.age, p2.marks, p2.isPass)

	fmt.Println()
}

type personInterface interface {
	student()
}

func callMethod(pI personInterface) {
	pI.student()
}

func main() {
	p1 := person1{
		name:   "Eddulamma",
		age:    31,
		marks:  45,
		isPass: true,
	}

	p2 := person2{
		name:   "Muddulamma",
		age:    32,
		marks:  65,
		isPass: true,
	}

	callMethod(p1)
	callMethod(p2)

}
