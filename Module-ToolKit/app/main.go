package main

import (
	"fmt"

	toolkit "github.com/KRISHNAHARI0618/gomodule"
)

func main() {
	var tools toolkit.Tools

	s := tools.RandomString(10)
	fmt.Println("RandomString : ", s)

}
