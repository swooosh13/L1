package main

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

/*
Реализовать постоянную запись данных в канал (главный поток).
Реализовать набор из N воркеров, которые читают произвольные
данные из канала и выводят в stdout.
Необходима возможность выбора количества воркеров при старте.

Программа должна завершаться по нажатию Ctrl+C.
Выбрать и обосновать способ завершения работы всех воркеров.

*/

func worker(wg *sync.WaitGroup, out <-chan []byte, done chan struct{}, i int) {
	defer wg.Done()
	for {
		select {
		case data, ok := <-out:
			if !ok {
				return
			}
			fmt.Fprintf(os.Stdout, "worker #%d, data: %s \n", i, data)
		case <-done: // срабатывает когда канал закрывается
			fmt.Fprintf(os.Stdout, "worker # %d has been stoped\n ", i)
			return
		}
	}
}

func main() {
	var N int
	fmt.Println("Scan N:")
	fmt.Scan(&N)

	in := make(chan []byte)
	done := make(chan struct{})

	var wg sync.WaitGroup
	for i := 0; i < N; i++ {
		wg.Add(1)
		go worker(&wg, in, done, i)
	}

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT)
	go func() {
		<-sig		// ждем сигнал SIGINT (ctrl + c)

		close(done) // корректно останавливаем все воркеры
		wg.Wait() 	//

		close(in) 	// завершаем поток ввода

		fmt.Println("Program has been finished with exit-code 0")
		os.Exit(0)
	}()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		in <- scanner.Bytes()
	}
}
