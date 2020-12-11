package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	go func(c chan int) {
		c <- 1
	}(c)
	x := <-c
	fmt.Println(x)
}
