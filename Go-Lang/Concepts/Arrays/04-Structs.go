package main

import "fmt"

func main() {
	fmt.Println("This is Structs Sections///")

	person1 := persons{
		name:    "peddireddy",
		age:     25,
		studied: false,
		marks:   75.4567,
	}
	fmt.Println(person1)

	// Anonymous Structs

	p := struct {
		name   string
		age    int
		number int
	}{
		name:   "peddireddy hari vardhan reddy",
		age:    45,
		number: 9493834674,
	}
	fmt.Println(p)

}

type persons struct {
	name    string
	age     int
	studied bool
	marks   float64
}
