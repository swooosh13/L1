package main

import (
	"fmt"
)

func SetBit(v *int64, pos int, t int) {
	if t == 1 {
		*v |= (1 << pos) // << 	- перемещает влево бит на кол-во позиций
	} else {
		*v &= ^(1 << pos) // ^ - xor operator (equal bit operator ~ - bitwise NOT)
	}
}

func main() {
	var v int64
	var i int
	fmt.Scan(&v, &i)

	// example
	// 5, 3
	// 101	- 5
	// 1101 - 13

	//
	// 5, 1
	// 101	- 5
	// 111  - 7

	//
	// 5, 0
	// 101	- 5
	// 111  - 5
	{
		v := v
		SetBit(&v ,i,1)
		fmt.Println(v)
	}

	// example
	// 5, 3
	// 101	- 5
	// 0101 - 5

	//
	// 5, 1
	// 101	- 5
	// 101  - 5

	//
	// 5, 0
	// 101	- 5
	// 100  - 4
	{
		v := v
		SetBit(&v ,i,0)
		fmt.Println(v)
	}
}
