package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

func main() {
	timeout := make(chan bool)
	end := make(chan bool)
	var iterator int
	var mutux sync.RWMutex
	var maxRoutineNum = 10
	ch := make(chan int, maxRoutineNum)
	go func() {
		time.Sleep(2 * time.Second)
		timeout <- true
	}()
	go func() {
		for i := 0; i < 30; i++ {
			ch <- 1
			go func(index int) {
				time.Sleep(10 * time.Second)
				fmt.Print(index)
				mutux.Lock()
				iterator++
				if iterator == 30 {
					end <- true
				}
				mutux.Unlock()
				<-ch
			}(i)
		}
	}()
	select {
	case <-timeout:
		log.Fatal("超时")
	case <-end:
		fmt.Println("成功")
	}
}
