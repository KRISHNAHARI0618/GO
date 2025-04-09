package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {

	fmt.Println("Hello World...")

	names := []string{"peddi", "reddy", "hari", "vardhan", "reddy"}

	for _, Val := range names {
		go greets(Val)
		wg.Add(1)
	}

	wg.Wait()

	defer fmt.Println("I will execute at the end....")
	panic fmt.Println("I am Panicing the orders....")

}

// Types Concurrency Routines Wat groups
// Case Insentive ; almost
// Variable type should be known in advance
// Everything is a Type

// Race Condition
// When mutiple go routiens are writing at a same time to one database
//

func greets(s string) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		fmt.Println(s)
	}
}
