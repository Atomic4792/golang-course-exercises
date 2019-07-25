package main

import "fmt"

type person struct {
	firstName, lastName string
	favFlavor           []string
}

func main() {
	p1 := person{
		firstName: "Tom",
		lastName:  "retriever",
		favFlavor: []string{"Vanilla", "Something"},
	}
	p2 := person{
		firstName: "Audi",
		lastName:  "retriever",
		favFlavor: []string{"Chocolate", "Something"},
	}
	fmt.Printf("Fist name: %v\nLast name: %v\nFavorite icecream flavors:", p1.firstName, p1.lastName)
	for _, v1 := range p1.favFlavor {
		fmt.Print(" ", v1)
	}
	fmt.Printf("\nFist name: %v\nLast name: %v\nFavorite icecream flavors:", p2.firstName, p2.lastName)
	for _, v2 := range p2.favFlavor {
		fmt.Print(" ", v2)
	}

}

/*
Create your own type “person” which will have an underlying type of “struct” so that it can store the following data:
first name
last name
favorite ice cream flavors
Create two VALUES of TYPE person. Print out the values, ranging over the elements in the slice which stores the favorite flavors

*/
