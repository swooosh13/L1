package main

import (
	"fmt"
	"math/rand"
	"time"
)

// It returns
// location of x in given array arr[l..r] is present,
// otherwise -1
func binarySearch(arr []int, l, r, x int) int {
	if r >= l {
		mid := l + (r-l)/2

		// If the element is present at the middle
		// itself
		if arr[mid] == x {
			return mid
		}

		// If element is smaller than mid, then
		// it can only be present in left subarray
		if arr[mid] > x {
			return binarySearch(arr, l, mid-1, x)
		}

		// Else the element can only be present
		// in right subarray
		return binarySearch(arr, mid+1, r, x)
	}

	// We reach here when element is not
	// present in array
	return -1
}

func partition(arr []int, low, high int) int {
	pivot := arr[high] // pivot
	i := low - 1		// Index of smaller element and indicates the right position of pivot found so far

	for j := low; j <= high-1; j++ {
		// if curr el is smaller than pivot
		if arr[j] < pivot {
			i++ 	// increment index of smaller element
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}

func qsort(nums []int, low, high int) {
	if low < high {
		/* pi is partitioning index, arr[p] is now
		   at right place */
		pi := partition(nums, low, high)

		qsort(nums, low, pi - 1)
		qsort(nums, pi + 1, high)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	max, low := 100, 1
	n := rand.Intn(max - low) + low
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = rand.Intn(100) + 1
	}

	qsort(arr, 0, len(arr) - 1)
	fmt.Println(arr)
	fmt.Println(binarySearch(arr, 0, len(arr) - 1, 42))
}
