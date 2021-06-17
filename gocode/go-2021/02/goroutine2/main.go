package main

import (
	"fmt"
	"sync"
)

const N = 5

func main() {
	m := make(map[int]int)
	// 使用waitgroup控制goroutine并发执行
	wg := &sync.WaitGroup{}
	//mu := &sync.Mutex{}
	wg.Add(N)
	for i := 0; i < N; i++ {
		go func(i int) {
			defer wg.Done()
			// 此处对map并发写加锁
			//mu.Lock()
			m[i] = i
			//mu.Unlock()
		}(i)
	}
	wg.Wait()
	// 遍历map元素
	for i, v := range m {
		fmt.Println(i, ":", v)
	}
}
