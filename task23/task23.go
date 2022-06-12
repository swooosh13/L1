package main

import "fmt"

func deleteFromIndex(nums []int, pos int) []int {
	if pos > len(nums) || pos < 0 {
		fmt.Println("Invalid index, out of range!")
	}

	return append(nums[0:pos], nums[pos + 1:]...)
}

func main() {
	arr := []int{1, 2, 5}
	arr = deleteFromIndex(arr, 1)
	fmt.Println(arr)
}
