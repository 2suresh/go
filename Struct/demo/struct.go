package main

import "fmt"

type Person struct {
	firstName string
	lastName  string
	Age       int
}

func main() {
	var p Person
	fmt.Println(p)

	p1 := Person{
		firstName: "Suresh",
		lastName:  "Kumar",
		Age:       35,
	}
	fmt.Println(p1)

	p2 := Person{"Parimala", "Deve", 38}
	fmt.Println(p2)

	p3 := Person{firstName: "Mallika"}
	fmt.Println(p3)
}
