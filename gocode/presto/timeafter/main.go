package main

import (
	"errors"
	"fmt"
	"time"
)

func main() {
	// ch := make(chan string)
	a := 1
	b := a
	a = 2
	fmt.Println(a, b)
}

func test() error {
	ch := make(chan string)
	go func() {
		time.Sleep(time.Second * 2)
		ch <- "result"
	}()
	select {
	case res := <-ch:
		fmt.Println(res)
	case <-time.After(time.Second * 1):
		fmt.Println("timeout")
		return errors.New("err")
	}
	return nil
}
