package main

import (
	"fmt"
)

func main() {

	a := 7 == 7
	b := 1 <= (1 * (0 + 1))
	c := 55 >= (1 << 5)
	d := 2 != 2
	e := (1024 >> 2) < 256
	f := 5*5 > 5*3
	fmt.Printf("%v\n%v\n%v\n%v\n%v\n%v\n", a, b, c, d, e, f)
}

/*
Using the following operators, write expressions and assign their values to variables:
==
<=
>=
!=
<
>
Now print each of the variables.

*/
