package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type counter uint64

func (c *counter) Inc() uint64 {
	return atomic.AddUint64((*uint64)(c), 1)
}

func (c *counter) Value() uint64 {
	return atomic.LoadUint64((*uint64)(c))
}

func (c *counter) Reset() uint64 {
	return atomic.SwapUint64((*uint64)(c), 0)
}

func main() {
	var c counter

	var wg sync.WaitGroup
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c.Inc()
		}()
	}

	wg.Wait()
	fmt.Println(c.Value())
}
