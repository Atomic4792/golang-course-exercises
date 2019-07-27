package main

import (
	"fmt"
)

func main() {

	var g func()

	f := func() {
		for i := 0; i < 3; i++ {
			fmt.Println(i)
		}

	}

	f()
	g = f
	g()
	fmt.Println("done")
}
