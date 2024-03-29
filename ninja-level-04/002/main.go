package main

import (
	"fmt"
)

func main() {
	mySlice := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	for _, v := range mySlice {
		fmt.Println(v)
	}
	fmt.Printf("%T", mySlice)
}

/*
Using a COMPOSITE LITERAL:
create a SLICE of TYPE int
assign 10 VALUES
Range over the slice and print the values out.
Using format printing
print out the TYPE of the slice

*/
