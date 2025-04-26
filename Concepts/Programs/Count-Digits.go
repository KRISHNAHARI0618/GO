package main

import "fmt"

func main() {
	fmt.Println("This is counting the numbers...")

	var num1 int
	fmt.Scanln(&num1)

	count := 0
	for num1 > 0 {
		num1 = num1 / 10
		count = count + 1
	}
	fmt.Printf("The count value is : %d \n ", count)
}
