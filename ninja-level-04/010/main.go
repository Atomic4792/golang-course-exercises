package main

import "fmt"

func main() {
	m := map[string][]string{
		`bond_james`:      []string{`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: []string{`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           []string{`Being evil`, `Ice cream`, `Sunsets`},
	}
	m[`Oscar`] = []string{`Coding`, `Soccer`, `Movies`}
	delete(m, `no_dr`)

	for i, v := range m {
		fmt.Printf("Record: %v\n", i)
		for j, v2 := range v { //v's used here, this just breaks down the values

			fmt.Println(j, v2)
		}

	}

}

/*
Using the code from the previous example, delete a record from your map. Now print the map out using the “range” loop
*/
