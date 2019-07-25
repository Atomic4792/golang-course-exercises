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
	/*
		fmt.Printf("Fist name: %v\nLast name: %v\nFavorite icecream flavors:", p1.firstName, p1.lastName)
		for _, v1 := range p1.favFlavor {
			fmt.Print(" ", v1)
		}
		fmt.Printf("\nFist name: %v\nLast name: %v\nFavorite icecream flavors:", p2.firstName, p2.lastName)
		for _, v2 := range p2.favFlavor {
			fmt.Print(" ", v2)
		}
	*/
	m := map[string]person{
		p1.lastName: p1,
		p2.lastName: p2,
	}

	for _, v := range m {
		fmt.Printf("%v\n", v.firstName)
		fmt.Printf("%v\n", v.lastName)

		for i, val := range v.favFlavor {
			fmt.Println(i, val)
		}
		fmt.Println("-------")
	}
	// I don't know why it doesn't print the two VALUES and the length of the map is 1 instead of two
}

/*
Take the code from the previous exercise, then store the values of type person in a map with the key of last name. Access each value in the map. Print out the values, ranging over the slice.
*/
