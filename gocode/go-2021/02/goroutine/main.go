package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

const N = 10

func main() {
	m := make(map[int]int)
	runtime.GOMAXPROCS(2)
	wg := &sync.WaitGroup{}
	mu := &sync.Mutex{}
	wg.Add(N)
	for i := 0; i < N; i++ {
		go func(i int) {
			defer wg.Done()
			mu.Lock()
			m[i] = i
			mu.Unlock()
		}(i)
	}
	wg.Wait()
	time.Sleep(time.Second)
	for i, v := range m {
		fmt.Println(i, ":", v)
	}
	println(len(m))
}
