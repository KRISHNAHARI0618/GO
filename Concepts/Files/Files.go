package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Helllo welcome to Go language....")
	content := "This should  i am ready to opened..."
	file, err := os.Create("hari.txt")
	if err != nil {
		panic(err)
	}
	length, err := io.WriteString(file, content)
	if err != nil {
		panic(err)
	}

	fmt.Println(length)
	readFile()
}

func readFile() {
	dataByte, err := os.ReadFile("hari.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(dataByte))
}

// Handling web Request
