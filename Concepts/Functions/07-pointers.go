package main

import "fmt"

func main() {

	var number int
	number = 45
	fmt.Println(number)

	var newNumber = &number
	fmt.Println(*newNumber)
	*newNumber = 58
	fmt.Println(*newNumber)
	fmt.Println(number)

	fmt.Print("Please Enter Number : ")
	var original int
	fmt.Scan(&original)
	fmt.Println("The number entered is : ", original)

	var addValue = &original
	fmt.Println(*addValue)

	*addValue = 100
	*addValue = original + 100
	fmt.Println(*addValue)

}
