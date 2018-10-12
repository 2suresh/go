package main

import "fmt"

type Car struct {
	Name, Modal, Color string
	weightInKG         float64
}

func main() {
	c := Car{
		Name:       "VW",
		Modal:      "Polo",
		Color:      "Blood Red",
		weightInKG: 2560,
	}

	d := &c
	fmt.Println((*d).weightInKG)
	fmt.Println(c.Name)
	fmt.Println(c.Modal)
	fmt.Println(c.Color)
}
