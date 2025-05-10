package main

import "fmt"

func main() {
	fmt.Println("Switch Case Statements for the functions .... ")

	num := 30
	switch {
	case num < 10:
		fmt.Printf("The Number is %d \n ", num)
	case num > 20:
		fmt.Printf("The number is %d \n ", num)
	default:
		fmt.Println("The New Language Number...")
	}

	num2 := 40

	switch {
	case num2 > 25:
		fmt.Printf("The Number is greater than %d \n ", num2)
	case num2 < 45:
		fmt.Printf("The Number is Less Than %d \n ", num2)
	default:
		fmt.Printf("The Numbers are equal %d \n", num2)
	}
}
