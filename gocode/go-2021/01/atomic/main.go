package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var value int32

func main() {
	var wg sync.WaitGroup
	wg.Add(3)
	for i := 1; i <= 3; i++ {
		go func(i int32) {
			defer wg.Done()
			addValue(i)
		}(int32(i))
	}
	wg.Wait()
	fmt.Println(value)

}

func addValue(delta int32) {
	for {
		v := value
		// 仅当CAS操作时才退出循环
		if atomic.CompareAndSwapInt32(&value, v, v+delta) {
			break
		}
	}
}
