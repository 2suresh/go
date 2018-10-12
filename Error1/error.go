package main

import (
	"errors"
	"fmt"
)

var errNotFound = errors.New("Item not found")

func getItem(a int) int {
	return a
}

func main() {
	// err := getItem(123)
	// fmt.Println(err)
	for i := 0; i < 5; i++ {
		fmt.Print(i, " \t")
	}
}
