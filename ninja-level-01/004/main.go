package main

import "fmt"

type jeff int

var x jeff

func main() {
	fmt.Printf("%v\t%T\n", x, x)
	x = 42
	fmt.Printf("%v", x)
}

/*
FYI - nice documentation and new terminology “underlying type”
https://golang.org/ref/spec#Types
For this exercise
Create your own type. Have the underlying type be an int.
create a VARIABLE of your new TYPE with the IDENTIFIER “x” using the “VAR” keyword
in func main
a)print out the value of the variable “x”
b)print out the type of the variable “x”
c)assign 42 to the VARIABLE “x” using the “=” OPERATOR
d)print out the value of the variable “x”

*/
