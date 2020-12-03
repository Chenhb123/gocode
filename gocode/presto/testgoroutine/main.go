package main

import (
	"fmt"
	"runtime"
	"sync"
)

/*
1.创建大量goroutine
2.每个goroutine不做任何事情，且阻塞不退出
3.统计创建前和创建后的内存消耗，计算平均值
*/
func main() {
	//runtime.GOMAXPROCS(3)

	memConsumed := func() uint64 {
		runtime.GC()
		var s runtime.MemStats
		runtime.ReadMemStats(&s)
		return s.Sys
	}

	var channel <-chan interface{}
	var wg sync.WaitGroup

	blockFunc := func() {
		wg.Done()
		<-channel
	}
	const numGoroutines = 1000000 // 1M
	wg.Add(numGoroutines)

	before := memConsumed()
	for i := numGoroutines; i > 0; i-- {
		go blockFunc()
	}
	wg.Wait()
	after := memConsumed()

	fmt.Printf("%.3fkb", float64(after-before)/numGoroutines/1024)
}
