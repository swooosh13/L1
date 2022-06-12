package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ReverseWords(s string) string {
	ss := strings.Split(s, " ")
	l := len(ss)

	for i := 0; i < l / 2; i++ {
		ss[i], ss[l - i - 1] = ss[l - i - 1], ss[i]
	}

	return strings.Join(ss, " ")
}

func main() {
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	text = strings.Trim(text, "\n")

	fmt.Println(ReverseWords(text))
}
