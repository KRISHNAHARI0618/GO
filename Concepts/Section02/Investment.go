package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Println("hello world .... ")

	const inflationRate float64 = 2.5
	var investAmount = 1000
	var expectRate = 5.5
	years := 10.0

	fmt.Println(investAmount)
	fmt.Println(expectRate)
	fmt.Println(years)

	var futureValue = float64(expectRate) * (1 + expectRate/100)
	fmt.Println(futureValue)

	var FutureValue = float64(expectRate) * math.Pow((1+expectRate/100), float64(years))
	fmt.Println(FutureValue)

	futureRealValue := futureValue / math.Pow((1+inflationRate/100), years)
	fmt.Print(futureRealValue, "\n")

	c, d := add(10, 20)

	fmt.Println(c)
	fmt.Println(d)

	fmt.Print("Please Enter Your Amount : ")
	fmt.Scan(&investAmount)
	fmt.Println("amount is ", investAmount)
	/*
		formatting the strings
	*/
	fmt.Printf("FutureValue %1.1v \n ", futureValue)
	FormateedStirng()

	f := switchCase('d')
	fmt.Println(f)
}

func add(a, b int) (int, int) {
	fmt.Println(a + b)
	a = a * 2
	b = b * 5
	if a < b {
		fmt.Println("Number is Larger::::")
	}
	return a, b
}

func FormateedStirng() {
	name := "Peddireddy"
	age := 25

	formattedString := fmt.Sprintf("My name is %s, and my age is %d ", name, age)
	fmt.Println(formattedString)
}

func switchCase(c byte) byte {
	switch {
	case '0' <= c && c <= '9':
		return c - '0'
	case 'a' <= c && c <= 'f':
		return c - 'a' + 10
	case 'A' <= c && c <= 'F':
		return c - 'A' + 10
	}
	return 0
}
func shouldEscape(c byte) bool{
	switch c {
	case '','?','&','+':
		return true
	}
	return false
}
