package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

/*

Дана последовательность чисел: 2,4,6,8,10.
Найти сумму их квадратов(22+32+42….) с использованием конкурентных вычислений.

*/

func SquareSum(nums []int) <-chan int {
	out := make(chan int)

	var sum int64
	go func() {
		var wg sync.WaitGroup
		for _, v := range nums {
			v := int64(v)
			wg.Add(1)
			go func(v int64) {
				defer wg.Done()
				atomic.AddInt64(&sum, v*v)
			}(v)
		}

		wg.Wait()
		out <- int(sum)
		close(out)
	}()

	return out
}

func main() {
	nums := []int{2, 4, 6, 8}
	sum := SquareSum(nums)
	fmt.Println(<-sum) // 120
}