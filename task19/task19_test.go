package main

import "testing"

//BenchmarkReverse-8       2267512               547.0 ns/op           288 B/op          2 allocs/op
//BenchmarkReverse1-8       590324              2473 ns/op            1616 B/op         50 allocs/op


func BenchmarkReverse(b *testing.B) {
	s := "The quick brown 狐 jumped over the lazy 犬 qwerty123"
	for i := 0; i < b.N; i++ {
		Reverse(s)
	}
}

func BenchmarkReverse1(b *testing.B) {
	s := "The quick brown 狐 jumped over the lazy 犬 qwerty123"
	for i := 0; i < b.N; i++ {
		Reverse1(s)
	}
}
