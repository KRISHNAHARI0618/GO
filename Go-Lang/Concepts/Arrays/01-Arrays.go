package main

import "fmt"

func main() {
	fmt.Println("This is Arrays Section...")

	a := [3]int{2, 3, 4}
	fmt.Println(a)

	b := [...]int{1, 2, 3, 4, 5, 6}
	fmt.Println(b)
	fmt.Println(len(b))

	var n []int = []int{1, 2, 3, 4}
	fmt.Println(n)

	var nums []int = []int{20, 30, 40, 50}
	fmt.Println(nums)

	for i := range nums {
		if nums[i] > 10 {
			fmt.Printf("The Number is : %v \n", nums[i])
		}
	}

	
}
