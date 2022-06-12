package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strings"
)

func main() {
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	text = strings.Trim(text, "\n")

	ss := strings.Split(text, " ")
	if len(ss) > 3 {
		fmt.Println("Too much args")
	}

	a, b := new(big.Int), new(big.Int)
	a.SetString(ss[0], 10)
	b.SetString(ss[2], 10)

	op := new(big.Int)
	switch ss[1] {
	case "+":
		op = op.Add(a, b)
	case "-":
		op = op.Sub(a, b)
	case "*":
		op = op.Mul(a, b)
	case "/":
		op = op.Div(a, b)
	default:
		fmt.Println("Unknown operation")
		return
	}
	
	fmt.Println(op.String())
}
