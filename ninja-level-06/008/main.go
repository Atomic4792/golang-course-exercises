package main

import "fmt"

func main() {
	xi := []int{11, 13, 5, 7, 8, 9, 10, 12, 15, 17, 18}
	e := even(smallest, xi...)
	o := odd(smallest, xi...)

	fmt.Printf("The smallest even number in this slice is: %v\n", e)
	fmt.Printf("The smallest odd number in this slice is: %v\n", o)

}

func smallest(x ...int) int {
	small := 1000
	for _, v := range x {
		if v < small {
			small = v
		}
	}
	return small
}

func even(f func(x ...int) int, j ...int) int {
	var yeet []int //will become a slice of even integers
	for _, v := range j {
		if v%2 == 0 {
			yeet = append(yeet, v)
		}
	}
	return f(yeet...)

}
func odd(f func(x ...int) int, k ...int) int {
	var yeet []int //will become a slice of odd integers
	for _, v := range k {
		if v%2 != 0 {
			yeet = append(yeet, v)
		}
	}
	return f(yeet...)

}

/*
A “callback” is when we pass a func into a func as an argument. For this exercise,
pass a func into a func as an argument
Remember, factions have roles/purpose

calculate the smallest even, odd and number in a slice
*/
