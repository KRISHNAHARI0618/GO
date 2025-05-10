package main

import "fmt"

func main() {
	veriadicFun(1, 2, 3, 4, 5)
	veriFunction(2, 3, 4, 5)
	// fmt.Println(s)
	summary := additionAll(2, 3, 4, 5, 66)
	fmt.Print(summary)

}

func veriadicFun(s ...int) {

	fmt.Println(s)
	n := 1
	for i := range len(s) {
		n = n + n*i
		fmt.Println(n)
	}
}

func veriFunction(a ...int) {
	fmt.Println(a)
}

func additionAll(x ...int) int {
	number := 1
	for v := range x {
		number = number + v*number
	}
	return number
}
