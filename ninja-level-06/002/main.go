package main

import "fmt"

func main() {

	s1 := foo([]int{10, 20, 30}...)
	fmt.Printf("foo() function returns a sum of: %v\n", s1)
	s2 := bar([]int{10, 20, 30})
	fmt.Printf("bar() function returns a sum of: %v\n", s2)
}

func foo(x ...int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}

func bar(j []int) int {
	sum := 0
	for _, v := range j {
		sum += v
	}
	return sum

}

/*
create a func with the identifier foo that
takes in a variadic parameter of type int
pass in a value of type []int into your func (unfurl the []int)
returns the sum of all values of type int passed in
create a func with the identifier bar that
takes in a parameter of type []int
returns the sum of all values of type int passed in

*/
