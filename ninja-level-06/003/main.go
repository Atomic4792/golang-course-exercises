package main

import "fmt"

func main() {
	defer funky()
	fmt.Println("I'm the fastest function here\n")

}

func funky() {
	defer func() {
		fmt.Println("I'm running slowly\n")
	}()
	fmt.Println("I'm running faster than the above function\n")
}

/*
Use the “defer” keyword to show that a deferred func runs after the func containing it exits.
*/
