package main

import "fmt"

func main() {
	fmt.Println("This is For Loops...")

	// Normal For Loop

	for i := 0; i < 10; i++ {
		fmt.Printf("The Number is %v \n", i)
	}

	// While Loop Syntax

	fmt.Println("This is Second Loops Functions")
	number := 1
	for number <= 20 {
		fmt.Printf("The number is %v \n", number)
		number++

	}

	simbaa := 1
	for {
		fmt.Printf("The Value is %v \n", simbaa)
		if simbaa >= 10 {
			break
		}
		simbaa++
	}

	// Printing Even Numbers

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			fmt.Printf("%v number is even number\n", i)
		}
	}

	// 
}
