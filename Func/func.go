package main

import "fmt"

func add(numA int, numB int) int {
	returnValue := (numA + numB)
	return returnValue
}

func main() {
	num_a, num_b := 25, 35
	sumValue := add(num_a, num_b)
	fmt.Println(sumValue)
}
