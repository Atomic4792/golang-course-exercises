package main

import (
	"fmt"
	"time"
)

func main() {
	i := time.Now().Year()
	born := 1996
	for born <= i {
		fmt.Println(born)
		born++
	}

}

/*
Create a for loop using this syntax
for condition { }
Have it print out the years you have been alive.

*/
