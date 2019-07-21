package main

import "fmt"

func main() {
	myArray := [5]int{}
	myArray[0] = 1
	myArray[1] = 2
	myArray[2] = 3
	myArray[3] = 4
	myArray[4] = 5

	for _, v := range myArray {
		fmt.Println(v)
	}

	fmt.Printf("\n%T", myArray)

}

/*
Using a COMPOSITE LITERAL:
create an ARRAY which holds 5 VALUES of TYPE int
assign VALUES to each index position.
Range over the array and print the values out.
Using format printing
print out the TYPE of the array

*/
