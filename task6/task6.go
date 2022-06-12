package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func g() {
	t := time.After(1 * time.Second)
	for {
		select {
		case <- t:
			fmt.Println("g4 exited after time")
		}
	}
}

// закрытие через done канал
func g1(done <-chan struct{}) {
	for {
		select {
		case _, ok := <-done:
			if !ok {
				fmt.Println("g1 exited")
				return
			}
			fmt.Println("g1 exited (done chan)")
			return
		}
	}
}

func g2(wg *sync.WaitGroup, i int) {
	defer wg.Done()

	time.Sleep(200 * time.Millisecond)
	fmt.Println("g2 exited,", i , "chan")
}


func g3(ctx context.Context) {
	for  {
		select {
		case <-ctx.Done():
			fmt.Println("g3 exited")
			return
		}
	}
}

func main() {
	done := make(chan struct{}, 2)

	go g()							// закроется сама после срабатывания собственного after
	time.Sleep(2 * time.Second)

	go g1(done)
	go g1(done)
	time.Sleep(1 * time.Second)
	done <- struct{}{} 				// посылаем в quit канал
	done <- struct{}{}
	time.Sleep(1 * time.Second)

	go g1(done)
	time.Sleep(1 * time.Second)
	close(done)						// завершение через закрытый канал
	time.Sleep(1 * time.Second)

	var wg sync.WaitGroup			// когда во всех горутинах сработает wg.Done()
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go g2(&wg, i)
	}
	wg.Wait()

	fmt.Println("with timeout")		// когда пройдет 1s все горутины завершатся,
	ctx, _ := context.WithTimeout(context.Background(), 1 * time.Second) // отменяется после определенного времени
	go g3(ctx)
	go g3(ctx)
	time.Sleep(2 * time.Second)

	fmt.Println("with deadline")		// когда пройдет 1s все горутины завершатся,
	ctx, _ = context.WithDeadline(context.Background(), time.Now().Add(1 * time.Second)) // отменяется после дедлайна (конкретное время) (now + 1s) или вызова функции отмены.
	go g3(ctx)
	go g3(ctx)
	time.Sleep(2 * time.Second)


	fmt.Println("with cancel") 		// когда вызовем cancel() все горутины завершатся
	ctx, cancel := context.WithCancel(context.Background())
	go g3(ctx)
	go g3(ctx)
	time.Sleep(1 * time.Second)			// имитируем работу программы до какого-то логического момента
	cancel()							// явное завершение

	time.Sleep(2 * time.Second)
}
