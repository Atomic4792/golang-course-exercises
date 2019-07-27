package main

import "fmt"

func main() {
	xi := []int{111, 4, 5, 6, 8, 9, 121, 2}

	odds := func(x ...int) int {
		sum := 0
		for _, v := range x {
			if v%3 == 0 {
				sum += v
			}
		}
		return sum
	}(xi...)

	low := func(j ...int) int {
		lowest := 1000
		//could've done it with range
		for i := 0; i < len(j); i++ {
			if j[i] < lowest {
				lowest = j[i]
			}

		}
		return lowest
	}(xi...)

	high := func(k ...int) int {
		highest := 0
		for _, v := range k {
			if v > highest {
				highest = v
			}
		}
		return highest
	}(xi...)

	fmt.Printf("The sum of odd numbers is: %v\n", odds)
	fmt.Printf("The lowest number is: %v\n", low)
	fmt.Printf("The highest number is: %v\n", high)

}

/*
Build and use an anonymous func

*/
