package main

import (
	"fmt"
)

var justString string

func createHugeString(n int) string {
	hugeStr := make([]rune, n)
	for i := 0; i < n; i++ {
		hugeStr[i] = 'ф'  	// 2 byte symbol
	}

	return string(hugeStr)
}

func someFunc() {
	v := createHugeString(1 << 12) // if peak (1 << 1)
	justString = v[:100]           // slice out of range [:100] with length 2
}

func someFuncV2() {
	v := createHugeString(1 << 4)
	rs := []rune(v)

	h := 100
	l := len(rs)

	if l >= h {
		justString = string(rs[:h])
		return
	}

	justString = v[:l]
}

func main() {
	//someFunc()
	someFuncV2()
	fmt.Println(justString)
}
