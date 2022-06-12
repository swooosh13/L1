package main

import (
	"fmt"
	"time"
)

func Sleep(d time.Duration) {
	<-time.After(d)
}

func main() {
	fmt.Println("start sleep")
	Sleep(10 * time.Second)
	fmt.Println("end")
}
