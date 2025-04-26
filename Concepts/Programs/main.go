package main

import (
	"fmt"
)

func main() {
	fmt.Println("Learning Go Programs..")

	var num1 int
	fmt.Printf("Please Enter num1 : ")
	fmt.Scanln(&num1)
	fmt.Println("Value is ..", num1)

	var num2 int
	fmt.Printf("Please Enter num2 : ")
	fmt.Scanln(&num2)
	fmt.Println("Value is ..", num2)
	fmt.Println(addition(num1, num2))

}

func addition(x, y int) int {
	return x + y
}
