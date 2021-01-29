package main

import (
	"fmt"
	"time"
)

/*
channel的哪些操作会引发panic
*/

func main() {
	// 1.关闭nil值channel
	//var ch chan struct{}
	//close(ch)

	// 2.关闭一个已关闭的channel
	//ch := make(chan struct{})
	//close(ch)
	//close(ch)

	// 3.向一个已经关闭的channel发送数据
	//ch := make(chan struct{})
	//close(ch)
	//ch <- struct{}{}

	// 从一个已经关闭的channel接收数据则不会panic
	ch := make(chan struct{})
	go func() {
		ch <- struct{}{}
		close(ch)
	}()
	fmt.Println(<-ch)
	time.Sleep(time.Second)
}
