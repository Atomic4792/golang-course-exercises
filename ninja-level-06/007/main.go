package main

import "fmt"

func main() {
	bar:=foo()
	bar()

}

func foo() func() {

	return func (){
		fmt.Println("I'm a function inside a function")}

}
/*
Create a func which returns a func
assign the returned func to a variable
call the returned func

 */
