package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)
	go func(c chan int) {
		c <- 1
	}(c)
	x := <-c
	fmt.Println(x)
	time.Sleep(10 * time.Second)
	fmt.Println("------")
	fmt.Println(<-c)
}
