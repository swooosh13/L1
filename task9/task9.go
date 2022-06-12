package main

import (
	"fmt"
	"os"
)

func receivePipe(nums []int, out chan<- int) {
	for _, v := range nums {
		out <- v
	}
	close(out)
}

func squarePipe(in <-chan int, out chan <- int) {
	for v := range in {
		out <- v*v
	}
	close(out)
}

func printPipe(in <-chan int) {
	for v := range in {
		fmt.Fprintf(os.Stdout, "%d ", v)
	}
}

func main()  {
	nums := []int{2, 4, 5, 10, 20}

	receiveCh := make(chan int) // в канал пишутся данные из массива
	sqCh := make(chan int)		// в канал записываются результаты x*x

	go receivePipe(nums, receiveCh)
	go squarePipe(receiveCh, sqCh)
	printPipe(sqCh)
}
