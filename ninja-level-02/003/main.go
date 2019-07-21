package main

import "fmt"

func main() {
	i := 256
	fmt.Printf("%d\t%b\t%#x\n", i, i, i)
	j := i << 1
	fmt.Printf("%d\t%b\t%#x", j, j, j)

}

/*
rite a program that
assigns an int to a variable
prints that int in decimal, binary, and hex
shifts the bits of that int over 1 position to the left, and assigns that to a variable
prints that variable in decimal, binary, and hex

*/
