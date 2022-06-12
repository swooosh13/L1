package main

import (
	"fmt"
)

func WhatType(v interface{}) {
	switch v.(type) {
	case int:
		fmt.Println("type int:", v)
	case string:
		fmt.Println("type string:", v)
	case bool:
		fmt.Println("type bool:", v)
	case chan interface{}:
		fmt.Println("type chan:", v, " - pointer")
	default:
		fmt.Println("unknown type")
	}
}

func main() {
	WhatType(1)
	WhatType("hello")
	WhatType(false)
	WhatType(make(chan interface{}))
	WhatType(9.99)
}
