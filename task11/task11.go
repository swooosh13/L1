package main

import "fmt"

func intersection(A, B []int) []int {
	inter := make([]int, 0)

	set := make(map[int]struct{}, len(A))
	for _, v := range A {
		set[v] = struct{}{}
	}

	for _, v := range B {
		if _, ok := set[v]; ok {
			inter = append(inter, v)
		}
	}

	return inter
}

func main() {
	A := []int{1, 3, 5, 0,- 1, 8}
	B := []int{0, 29, 8, 2}

	fmt.Println(intersection(A, B))
}
