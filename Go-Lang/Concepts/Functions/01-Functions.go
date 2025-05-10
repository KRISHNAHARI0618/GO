package main

import "fmt"

func main() {
	fmt.Println("This is Functions Class...")

	Method1()
	Method2("Peddireddy")
	s1 := Method3("Peddireddy")
	fmt.Println(s1)

	s2, s3 := Method4("peddirewddy")
	fmt.Println(s2, s3)
}

func Method1() {
	fmt.Println("This is First Method With No Returns and No Parameters...")
}

func Method2(s string) {
	fmt.Println("Calling Second Method2 ...")
	fmt.Println("My Name is ", s)
}

func Method3(s string) (s1 string) {
	s1 = s[2:]
	return s1
}

func Method4(s string) (age int, name string) {
	age = len(s)
	name = s[3:7]
	return age, name
}
