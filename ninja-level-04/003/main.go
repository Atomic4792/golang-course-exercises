package main

import (
	"fmt"
)

func main() {
	mySlice := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	for _, v := range mySlice {
		fmt.Println(v)
	}
	fmt.Printf("%T\n", mySlice)

	newSlice1 := mySlice[:5]
	fmt.Println(newSlice1)
	newSlice2 := mySlice[5:]
	fmt.Println(newSlice2)
	newSlice3 := mySlice[2:7]
	fmt.Println(newSlice3)
	newSlice4 := mySlice[1:6]
	fmt.Println(newSlice4)


}

/*
Using the code from the previous example, use SLICING to create the following new slices which are then printed:
[42 43 44 45 46]
[47 48 49 50 51]
[44 45 46 47 48]
[43 44 45 46 47]

*/
