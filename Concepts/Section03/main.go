package main

import (
	"bufio"
	"fmt"
	"os"
)

const Name = "peddireddy hari vardhan reddy"

var id = 345

func main() {
	var username string = "peddireddy"
	fmt.Printf("Variables is of Type : %T", username)
	fmt.Println()

	var users string
	fmt.Printf("Variables is of Type : %T", users)
	fmt.Println(users)

	numebrs := 300000
	fmt.Println(numebrs)

	var num = 40000
	fmt.Println(num)
	add()

	reader := bufio.NewReader(os.Stdin)
	fmt.Println(reader)

	input, _ := reader.ReadString('\n')
	fmt.Println("You have entered as : ", input)

}

func add() {
	fmt.Println(Name)
	fmt.Println(id)
}
