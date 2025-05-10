package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("This is New Added Life ... ")
	sum()
	forLoop()
}

func sum() {
	fmt.Println("added new life")

	x := 0
	fmt.Scanf("%d", &x)
	fmt.Printf("The Number is %v \n ", x)

	if z := 2 * rand.Intn(x); z <= x {
		fmt.Printf("z is %v and x is %v \n ", z, x)
	} else {
		fmt.Println("Added New Lifes ")
	}

	num1 := 0
	fmt.Scanf("%d", &num1)

	num2 := 0
	fmt.Scanf("%d", &num2)

	// if x >= numberAdd(num1, num2) {
	// 	fmt.Printf("num1 is %v and num2 is %v \n ", num1, num2)
	// } else {
	// 	fmt.Println("Sorry Number is wrong")

	// }

	if num, err := numberAdd(num1, num2); err == true {
		fmt.Println("Number is ", num)
	}

}

func numberAdd(a, b int) (int, bool) {
	fmt.Println("Adding the Numbers ... ")
	sum := a + b
	return sum, true

}

func forLoop() {
	fmt.Println("This is For Loop Implementation...")
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}
	for nums := range 10 {
		fmt.Println(nums)
	}

	day := 3

	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednessday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	default:
		fmt.Println("Sorry No Day Found")
	}
}
