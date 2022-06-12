package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Reverse(s string) string {
	// Get Unicode code points.
	n := 0
	rs := make([]rune, len(s))
	// r - rune type
	for _, r := range s {
		rs[n] = r
		n++
	}
	rs = rs[0:n]

	for i := 0; i < n/2; i++ {
		rs[i], rs[n-1-i] = rs[n-1-i], rs[i]
	}

	return string(rs)
}

func Reverse1(s string) string {
	var result string
	for _, v := range s {
		result = string(v) + result
	}
	return result
}

func main() {
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	text = strings.Trim(text, "\n")

	fmt.Println(Reverse(text))
}
