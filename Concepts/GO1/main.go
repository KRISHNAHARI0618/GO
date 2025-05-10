package main

import "fmt"

func main() {
	fmt.Println("Hello world...")
	a := 42
	fmt.Println("This is a value :", a)
	a, b, c, d := 1, "peddireddy", true, 34.67
	fmt.Println(a, b, c, d)

	var e int = 20
	fmt.Printf("This is Varibale e %d", e)
	fmt.Println()

	a1, b1 := 10, 20
	fmt.Printf("%v \t %b \t %x \n ", a1, a1, a1)
	fmt.Printf("%v \t %b \t %x \n ", b1, b1, b1)

}
