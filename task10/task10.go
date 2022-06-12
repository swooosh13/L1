package main

import "fmt"

func main() {
	nums := []float64{25.4, -27.0, 13.0, 19.0,
		15.5, 24.5, -21.0, 32.5}

	temp := make(map[int][]float64)
	for _, v := range nums {
		k := int(v) - int(v) % 10
		temp[k] = append(temp[k], v)
	}

	fmt.Println(temp)
}
