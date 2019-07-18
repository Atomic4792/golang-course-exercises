package main

import (
	"fmt"
	"time"
)

func main() {
	born:=1996
	thisYear:=time.Now().Year()
	for {
		if born>thisYear{
			break
		}
		fmt.Println(born)
		born++
	}

}
/*
Create a for loop using this syntax
for { }
Have it print out the years you have been alive.

 */