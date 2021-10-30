package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	/*
		alex := person{
			firstName: "Alex",
			lastName:  "Dunphy",
		}
	*/
	var alex person
	alex.firstName = "Alex"
	alex.lastName = "Dunphy"

	fmt.Println(alex)
}
