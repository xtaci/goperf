package main

import (
	"container/list"
	"testing"
)

func BenchmarkLoopSlice(b *testing.B) {
	s := make([]byte, b.N)
	b.ResetTimer()

	for range s {
	}
}

func BenchmarkLoopList(b *testing.B) {
	l := list.New()
	for i := 0; i < b.N; i++ {
		l.PushBack(nil)
	}

	b.ResetTimer()
	for e := l.Front(); e != nil; e = e.Next() {
	}
}
