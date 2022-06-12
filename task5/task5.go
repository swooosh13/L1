package main

import (
	"fmt"
	"math/rand"
	"time"
)

func receiver(ch chan int) {
	for {
		select {
		case x, ok := <-ch:
			if !ok {
				return
			}
			fmt.Println(x)
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	var N int
	fmt.Println("Duration (seconds):")
	fmt.Scan(&N)

	ch := make(chan int)

	go receiver(ch)

	t := time.After(time.Duration(N) * time.Second)
	for {
		select {
		case <-t:  // N секунд сработает
			close(ch)  	// закрываем канал и завершаем программу
			return
		default:
			ch <- rand.Intn(100)		// записываем в канал с интервалом 200 миллисекунд
			time.Sleep(200 * time.Millisecond)
		}
	}
}
