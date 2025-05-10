package main

import "fmt"

func main() {
	fmt.Println("Make Functions....")
	num := make([]int, 0, 10)
	fmt.Println(num)
	fmt.Println(len(num))
	fmt.Println(cap(num))

	// Multidimensional Slices [] []

	num = append(num, 0, 1, 2, 3, 5, 6, 4, 6, 7, 8, 3, 2, 1, 3, 5)
	fmt.Println(num)

}
