package main

import (
	"fmt"
)

func main() {
	fmt.Println("Slices Function ")

	var num []int = []int{20, 30, 40, 50}
	fmt.Println(num)

	for i := 0; i < len(num); i++ {
		fmt.Println(num[i])
	}

	var a []int

	for _, v := range num {

		if v%2 == 0 {
			a = append(a, v*2)
		}

	}
	fmt.Println(a)

	count := 0
	// Veriadic Function
	// [...] any number of values we can give in variadic functions

	for {
		if count <= 10 {
			fmt.Println(count)
			count++
		} else {
			break
		}
	}

	var nums []int = []int{20, 30, 40, 50, 60}
	fmt.Println(nums[0:3])

	nums = append(nums[:1], nums[3:]...)
	fmt.Println(nums)
}
