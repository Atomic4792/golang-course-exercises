package main

import "fmt"

func main() {
	x := foo()
	fmt.Printf("My lucky number is: %v\n", x)
	a, b := bar()
	fmt.Println(b, a)

}
func foo() int {
	return 52

}
func bar() (int, string) {
	i := 1996
	j := "I was born in :"

	return i, j
}


/*
create a func with the identifier foo that returns an int
create a func with the identifier bar that returns an int and a string
call both funcs
print out their results

*/
