package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"time"
)

func main() {
	var name = "peddireddy"
	fmt.Println(name)
	fmt.Println("Pedireddy Hari Vardhan Reddy...")

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the text : ")
	str, _ := reader.ReadString('\n')
	fmt.Println(str, '\n')
	i1, i2, i3 := 12, 23, 76
	sumInt := i1 + i2 + i3
	fmt.Println(sumInt)

	for i := 0; i <= 10; i++ {
		fmt.Printf("Iteration : %v\n", i)
	}
	count := 0
	for count < 3 {
		fmt.Printf("Iteration: %v\n", count)
		count++
	}

	n := time.Now()
	fmt.Printf("The current time is %s", n)

	fmt.Printf(n.Format(time.ANSIC))

	fmt.Println()
	var colors = make([]string, 0, 4)
	colors = append(colors, "red", "green", "yellow", "pink")

	fmt.Println(colors)

	number := 64
	var IntNumber = &number
	*IntNumber = *IntNumber + 100
	fmt.Println(*IntNumber)
	fmt.Println(number)

	var numbers [3]int
	numbers[0] = 1
	numbers[1] = 2
	numbers[2] = 3
	fmt.Println(numbers)

	var nums = [5]int{1, 2, 3, 4, 5}

	fmt.Println(nums)
	fmt.Println(len(nums))

	var num = []int{7, 8, 9, 4, 2, 5, 8, 9, 3, 2}
	fmt.Println(num)

	sort.Ints(num)
	fmt.Println(num)

	fmt.Printf("adding %v and %v gives %v\n", 20, 30, 50)

}
