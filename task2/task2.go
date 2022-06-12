package main

import (
	"fmt"
	"sync"
	"time"
)

// synch
func Square(nums []int) <-chan int {
	out := make(chan int, len(nums))

	go func() {
		start := time.Now()
		for _, v := range nums {
			time.Sleep(1 * time.Second)
			out <- v * v
		}

		fmt.Println(time.Since(start))
		close(out)
	}()

	return out
}

// concurrent
func SquareV2(nums []int) <-chan int {
	out := make(chan int, len(nums))

	go func() {
		start := time.Now()

		var wg sync.WaitGroup
		for _, v := range nums {
			wg.Add(1)
			go func(v int) {
				defer wg.Done()
				time.Sleep(1 * time.Second)
				out <- v * v
			}(v)
		}
		wg.Wait()

		fmt.Println(time.Since(start))
		close(out)
	}()

	return out
}

func main() {
	nums := []int{2, 4, 6, 8, 10}

	out := Square(nums) // 5s
	for v := range out {
		fmt.Println(v)
	}
	// 4 16 36 64

	out2 := SquareV2(nums) // 1s
	for v := range out2 {
		fmt.Println(v)
	}
	// 4 100 36 64 16
}
