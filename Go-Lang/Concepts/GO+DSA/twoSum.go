package main

import "fmt"

func main() {
	fmt.Println("This is Two Sum Array....")

	list := []int{2, 3, 4, 5, 6}

	for i, val := range list {
		fmt.Println(i, val)
	}
	result := TwoSumArray(list, 9)
	fmt.Println(result)

}

func TwoSumArray(arr []int, target int) []int {

	numMap := make(map[int]int)

	for i, val := range arr {
		numMap[i] = val
	}
	fmt.Println(numMap)

	for i, val := range arr {
		compliment := target - val
		if index, found := numMap[compliment]; found {
			return []int{index, i}
		}
		numMap[val] = i
	}
	return nil
}
