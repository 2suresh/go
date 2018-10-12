package main

import "fmt"

func count(i int) (n int) {

	defer func(i int) {
		n = n + i
	}(i)

	i = i * 2
	n = i
	return
}

func main() {
	fmt.Println(count(5))
}
