package main

import "fmt"

type person struct {
	first, last string
	age         int
}

func (p person) speak() {
	fmt.Printf("Hi, my name is %v and I'm %v years old", p.first, p.age)
}
func main() {
	p1 := person{
		first: "Steve",
		last:  "Jobs",
		age:   51,
	}
	p1.speak()
}

/*
Create a user defined struct with
the identifier “person”
the fields:
first
last
age
attach a method to type person with
the identifier “speak”
the method should have the person say their name and age
create a value of type person
call the method from the value of type person

*/
