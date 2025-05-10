package main

import (
	"fmt"
)

func init() {
	fmt.Println("Initialization Code Running ...")
}

func main() {
	fmt.Println("Peddireddy Hari Vardhan Reddy")
	conditionals()
	addition()
}

func conditionals() {
	fmt.Println("Learning Conditionals In Go Language...")

	fmt.Print("Please Enter The Number : ")

	var number int
	fmt.Scanf("%d", &number)
	fmt.Println("The number is : ", number)
}

func addition() {
	var number int
	fmt.Print("Please Enter the number : ")
	fmt.Scanf("%d", &number)
	if number > 20 {
		fmt.Println("Number is Greater Than Expected ..")
	} else {
		fmt.Println("The number is Less Than Expected ..")
	}
}
