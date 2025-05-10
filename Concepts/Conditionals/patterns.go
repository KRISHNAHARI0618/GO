package main

import "fmt"

func main() {
	fmt.Println(" Printing Patterns Problems ")

	// Right Angled Triangle

	for i := 0; i <= 5; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	// Left Angled Triangle

	for i := 5; i >= 0; i-- {
		for j := 0; j <= i; j++ {
			fmt.Printf("%v", j)
		}
		fmt.Println()
	}

	
}
