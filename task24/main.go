package main

import (
	"fmt"
	"l1/task24/point"
)


func main() {
	a, b := point.NewPoint(0, 0), point.NewPoint(3.0, 4.0)
	fmt.Println(point.Distance(a, b)) // 5 (египетский треугольник)
}