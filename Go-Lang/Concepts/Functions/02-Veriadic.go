package main

import "fmt"

func main() {
	veriadicFun(1, 2, 3, 4, 5)
	// fmt.Println(s)

}

func veriadicFun(s ...int) {

	fmt.Println(s)
	n := 1
	for i := range len(s) {
		n = n + n*i
		fmt.Println(n)
	}
}
