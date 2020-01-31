package main

import (
	"sync"
	"sync/atomic"
	"testing"
)

func BenchmarkMutex(b *testing.B) {
	var mu sync.Mutex
	for i := 0; i < b.N; i++ {
		mu.Lock()
		mu.Unlock()
	}
}

func BenchmarkRWMutexR(b *testing.B) {
	var mu sync.RWMutex
	for i := 0; i < b.N; i++ {
		mu.RLock()
		mu.RUnlock()
	}
}

func BenchmarkRWMutexW(b *testing.B) {
	var mu sync.RWMutex
	for i := 0; i < b.N; i++ {
		mu.Lock()
		mu.Unlock()
	}
}

func BenchmarkAtomic(b *testing.B) {
	var sum int64
	for i := 0; i < b.N; i++ {
		atomic.AddInt64(&sum, int64(i))
	}
}
