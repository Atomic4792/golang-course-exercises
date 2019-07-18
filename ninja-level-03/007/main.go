package main

import "fmt"

func main() {
	favSport := "Soccer"
	switch favSport {
	case "Tennis":
		fmt.Println("Federer is the best hands down")
	case "Soccer":
		fmt.Println("CR7 for life")

	case "Basketball":
		fmt.Println("Go Lakers")

	default:
		fmt.Println("I luv sports")

	}
}

/*
Create a program that uses a switch statement with the switch expression specified as a variable of TYPE string with the IDENTIFIER “favSport”.
*/

/*
This is ridiculous without user input
*/
