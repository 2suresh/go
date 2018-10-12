package main

import "fmt"

func finished() {
	fmt.Println("finishing find the largested")
}

func findLarge(nums []int) {
	defer finished()
	fmt.Println("Strating find the largested")
	fmt.Println(nums)
	max := nums[0]
	for _, v := range nums {
		if v > max {
			max = v
		}
	}
	fmt.Println("Largested number is ", max)
}

func main() {
	nums := []int{23, 55, 20, 19, 47}
	findLarge(nums)
}
