package main

import "fmt"

func main() {

	anon := struct {
		name      string
		luckyNums []int
		favDay    map[string]string
	}{
		name:      "Ethan",
		luckyNums: []int{1, 2, 5},
		favDay: map[string]string{
			"Ethan": "Monday",
			"Hila":  "Friday",
		},
	}

	fmt.Printf("Name: %v\nLucky Numbers: %v\nFavorite day: %v\n", anon.name, anon.luckyNums, anon.favDay["Ethan"])
	for k, v := range anon.favDay {
		fmt.Println(k, v)
	}

}
/*
Create and use an anonymous struct.

 */