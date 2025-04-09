package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	inputValue := bufio.NewReader(os.Stdin)

	input, err := inputValue.ReadString('\n')
	fmt.Println(input)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(input)
	}

	timeSet()

}

// Handling Time in Go
func timeSet() {
	fmt.Println("Welcome To Time Study....")
	presentTime := time.Now()
	fmt.Println(presentTime)
	fmt.Println(presentTime.Format("01-02-2006"))
	createdDate := time.Date(2025,time.Month(9),20,15,15,15,15,time.Now().UTC().Location())
	fmt.Println(createdDate)
}
