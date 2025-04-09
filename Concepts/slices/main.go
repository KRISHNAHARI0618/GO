package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	highScores := make([]int, 10)
	fmt.Println(highScores)
	highScores[1] = 20
	highScores[2] = 50
	highScores[3] = 60
	highScores[6] = 80
	highScores[5] = 67
	highScores[7] = 89

	highScores = append(highScores, 10, 20, 304, 05, 06, 07, 46, 859)
	fmt.Println(highScores)
	sort.Ints(highScores)
	fmt.Println(highScores)

	//Ignoring the values:

	index := 5
	newIndexValues := append(highScores[:index], highScores[index+1:]...)
	fmt.Println(newIndexValues)

	// Maps are like key value pairs format key can be anything and value can be anything
	// make(map[string]string)
	// langugaes[] = "javascript"
	maps := make(map[string]string)
	fmt.Println(maps)
	maps["js"] = "javascript"
	maps["py"] = "python"
	maps["sc"] = "science"
	fmt.Println(maps)

	delete(maps, "js")
	fmt.Println(maps)

	// Loops are intresting in Go Lang

	for key, value := range maps {
		fmt.Println(key, ":", value)
	}

	userDetails := User{"peddireddy", 25}
	fmt.Println(userDetails)

	// Control Flow strutures

	if true {
		fmt.Println("Peddireddy Hari Vardhan Reddy")
	} else {
		fmt.Println("Calling another time ")
	}

	student1 := DateApp{19, "peddireddy", "kalasalingam", 2021}
	fmt.Println(student1)

	if student1.Age > 20 {
		fmt.Println("Please Login ....")
	} else {
		fmt.Println("First Get out from app and study well .......")
	}

	rand.Seed(time.Now().Unix())
	numberNow := rand.Intn(100)
	fmt.Println(numberNow)

	days := []int{10, 20, 40, 50, 80, 35, 76}
	for i := 0; i < len(days); i++ {
		fmt.Println(days[i])
	}
	value := 0
	for value < 10 {

		if value == 7 {
			goto peddi
		}

		if value == 5 {
			break
		}

		if value == 6 {
			value++
			continue
		}

		fmt.Println(value)
		value++

	}

peddi:
	fmt.Println("Printing the Go To Statement in Years....")

	for val := range 10 {
		fmt.Print(val, " ")
	}
	fmt.Println()
	addition()

	student1.getstatus()

	nums := sum(10, 20, 30, 40, 50, 60, 70, 80, 90, 100)
	fmt.Println(nums)

	rever()

	mydefer()
}

type User struct {
	Name string
	Age  int
}

type DateApp struct {
	Age       int
	Name      string
	College   string
	PassedOut int
}

// Functions are more useful when you are learning any language

func addition() {
	fmt.Println("Addition of two numbers ..... ")
}

func (da DateApp) getstatus() {

	fmt.Println(da.Name)
}

func sum(nums ...int) int {

	total := 0
	for _, n := range nums {
		total = total + n
	}
	return total
}

func rever() {
	for i := 10; i > 0; i-- {
		fmt.Print(i, " ")
	}
	fmt.Println()
}

func mydefer() {
	defer fmt.Println("World in Go ...")
	defer fmt.Println("Adding to Queue...")
	fmt.Println("Hello ...")
}

// Methods : passing structs to the function is called the method
// To work with any other files we require libraries
