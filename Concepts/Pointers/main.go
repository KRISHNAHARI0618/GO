// Pointers advantages
// Avoid unnecesary value copies
// Direct Mutate Values
// Closures: in ananymous function when we use normal funcion values is called closures..
package main

import "fmt"

type transafm func(int) int

type Person struct {
	Name string
	age  int
}

func main() {

	person1 := Person{Name: "hari", age: 25}
	fmt.Println(person1)

	nums := [5]int{1, 2, 3, 4, 5}
	fmt.Println(len(nums))

	slice2 := make([]int, 5)
	fmt.Println(append(slice2, 6))

	age := 32
	fmt.Println(age)

	var agePointer *int = &age
	fmt.Println(agePointer)
	fmt.Println(*agePointer)

	var number int
	var name string
	var percent float64
	var commited bool

	fmt.Println(number, name, percent, commited)
	adultYears := ageDeduct(agePointer)
	fmt.Println(adultYears)

	editValue(agePointer)
	fmt.Println(age)
	listNums := []int{1, 2, 3, 4}
	dble := transFromNumbers(&listNums, double)
	fmt.Println(dble)
	trple := transFromNumbers(&listNums, triple)
	fmt.Println(trple)

	DeferFunc()

	num := 5
	fmt.Println(Recur(num))

	fmt.Println(variadic(20,30,40,50))

	greet := func(name string) {
		fmt.Println("Hello ", name)
	}
	greet("Peddireddy")

}

func ageDeduct(age *int) int {
	return *age - 18
}

func editValue(age *int) {
	*age = *age - 18
}

// Functions Deep Dive
// using Functions as values
// Ananomous Functions
// Recursions
// Variadic Fucntions

func transFromNumbers(numbers *[]int, transafm func(int) int) []int {

	dnumbers := []int{}
	for _, val := range *numbers {
		dnumbers = append(dnumbers, transafm(val))
	}
	return dnumbers

}

func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}

func DeferFunc(){
	defer fmt.Println("Good Byee Dear")
	fmt.Println("Hwllo Mr....")
}

func Recur(num int)int{
	if num == 0{
		return 1
	}
	return num * Recur(num -1)
}

func variadic(nums ...int) int{
	total := 0
	for _,n := range nums{
		total = total + n
	}
	return total
}


type speaker interface{
	speak() string
}
type dog struct{}

func (d dog)speak() string{
	return "Woof"
}
