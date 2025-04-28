package main

import (
	"fmt"
	"mymodule/local-modules/simplemodule"
)

func main() {
	fmt.Println("Adding Two numbers...")
	result := simplemodule.Multiple(20, 30)
	fmt.Println(result)
}
