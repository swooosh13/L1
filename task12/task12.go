package main

import "fmt"

type Set map[string]struct{}

func main() {
	str := []string{"cat", "cat", "dog","cat", "tree"}

	set := Set{}
	for _, v := range str {
		set[v] = struct{}{}
	}

	for k, _ := range set {
		fmt.Println(k)
	}

	fmt.Println(set)
	delete(set, "cat")
	fmt.Println(set)
}
