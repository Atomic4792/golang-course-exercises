package main

import "fmt"

func main() {
	normieSlice1 := []string{"James", "Bond", "Shaken, not stirred"}
	normieSlice2 := []string{"Miss", "Moneypenny", "Helloooooo, James."}
	multiSlice := [][]string{normieSlice1, normieSlice2}

	for i, v := range multiSlice {
		fmt.Printf("Position: %v Value: %v\n", i, v)

	}
	for i := 0; i < len(multiSlice); i++ {
		fmt.Printf("Record: %v\n", i)
		for j, v := range multiSlice[i] {
			fmt.Printf("Position: %v \tValue: %v \t\n", j, v)

		}

	}

}

/*
Create a slice of a slice of string ([][]string). Store the following data in the multi-dimensional slice:

"James", "Bond", "Shaken, not stirred"
"Miss", "Moneypenny", "Helloooooo, James."

Range over the records, then range over the data in each record.

*/
