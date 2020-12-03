package main

import (
	"fmt"
	"sync"
)

func main() {
	// test1()
	test2()
}

func test2() {
	a := make([]int, 0)
	ch := make(chan int)
	var wg sync.WaitGroup
	go func() {
		for v := range ch {
			a = append(a, v)
		}
	}()
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func(i int) {
			ch <- i
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println(len(a))
}

func test1() {
	a := make([]int, 0)
	var wg sync.WaitGroup
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func(i int) {
			a = append(a, i)
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println(len(a))
}
