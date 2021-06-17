package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	ch := make(chan int, 1024)
	go func(ch chan int) {
		for {
			val := <-ch
			fmt.Printf("val:%d\n", val)
		}
	}(ch)

	tick := time.NewTicker(1 * time.Second)
	for i := 0; i < 5; i++ {
		select {
		case ch <- i:
		case <-tick.C:
			fmt.Printf("%d: case <-tick.C\n", i)
		}

		time.Sleep(500 * time.Millisecond)
	}
	close(ch)
	tick.Stop()
	m := sync.Map{}
}
