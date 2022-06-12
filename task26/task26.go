package main

import (
	"fmt"
	"strings"
)

func IsUnique(str string) bool {
	s := strings.ToLower(str)
	set := make(map[rune]struct{})

	for _, v := range s {
		if _, ok := set[v]; ok {
			return false
		}
		set[v] = struct{}{}
	}

	return true
}

func main() {
	fmt.Println(IsUnique("狐ju"))     // true
	fmt.Println(IsUnique("狐ju狐"))    // false
	fmt.Println(IsUnique("abcd"))    // true
	fmt.Println(IsUnique("abcDeFA")) // false
	fmt.Println(IsUnique("aabcd"))   // false
}
