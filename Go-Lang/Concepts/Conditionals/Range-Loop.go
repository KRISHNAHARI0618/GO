package main

import "fmt"

func main() {
	fmt.Println("This is Range Loop")

	// Data Structures Slices
	x := []int{20, 30, 40, 50, 60, 70, 80}
	for i, j := range x {
		fmt.Println(i, j)
	}

	// Data structures - Map

	map1 := map[string]int{

		"name": 43,
		"age":  25,
	}
	fmt.Println(map1)
	for k, v := range map1 {
		fmt.Println(k, v)
	}

	//This is map data structures

	m1 := map[string]int{}
	m1["name"] = 1234
	m1["age"] = 23
	m1["number"] = 9493834674
	fmt.Println(m1)

	// This is slice Data Structures
	x2 := []int{}
	x2 = append(x2, 20, 30, 40, 50, 60)
	fmt.Println(x2)

	// Array With Fixed Length
	x3 := [5]int{10, 100, 1000, 10000, 100000}
	fmt.Println(x3)

	// Finding remainder and Modulus

	x5 := 10
	y2 := 3

	c1 := x5 / y2
	c2 := x5 % y2

	fmt.Println(c1, c2)

}
