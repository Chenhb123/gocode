package test

import (
	"sync"
	"testing"
)

var mu sync.Mutex

// Call 。
func Call() {
	mu.Lock()
	mu.Unlock()
}

// CallDefer .
func CallDefer() {
	mu.Lock()
	defer mu.Unlock()
}

// BenchmarkCall 。
func BenchmarkCall(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Call()
	}
}

// BenchmarkDefer .
func BenchmarkDefer(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CallDefer()
	}
}
