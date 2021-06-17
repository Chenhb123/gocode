package slower

import (
	"runtime"
	"sync"
)

const limit = 10000000000

func SerialSum() int {
	sum := 0
	for i := 0; i < limit; i++ {
		sum += i
	}
	return sum
}

func ConcurrentSum() int {
	n := runtime.GOMAXPROCS(0)
	sums := make([]int, n)
	wg := sync.WaitGroup{}
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			start := (limit / n) * i
			end := start + limit/n
			for j := start; j < end; j++ {
				sums[i] += j
			}
		}(i)
	}
	wg.Wait()
	sum := 0
	for _, v := range sums {
		sum += v
	}
	return sum
}

func ChannelSum() int {
	ch := make(chan int)
	n := runtime.GOMAXPROCS(0)
	for i := 0; i < n; i++ {
		go func(i int, ch chan<- int) {
			sum := 0
			start := (limit / n) * i
			end := start + limit/n
			for j := start; j < end; j++ {
				sum += j
			}
			ch <- sum
		}(i, ch)
	}
	sum := 0
	for i := 0; i < n; i++ {
		sum += <-ch
	}
	return sum
}
