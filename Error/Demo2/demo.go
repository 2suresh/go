package main

import (
	"fmt"
	"os"
)

func main() {
	// a, err := os.Open("texts.txt")
	// if err, ok := err.(*os.PathError); ok {
	// 	fmt.Println("File at path", err.Path, "failed to open")
	// 	return
	// }
	// fmt.Printf("%+v\n", &a)
	// fmt.Println(a.Name(), "File opened successfully")
	a, err := os.Open("texts.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(a.Name())
	fmt.Println(a.FileInfo.Size())
}
