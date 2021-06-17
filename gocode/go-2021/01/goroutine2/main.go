package main

import (
	"fmt"
	"time"
)

func main() {
	chs := make([]chan struct{}, 4)
	length := len(chs)
	for i := 0; i < length; i++ {
		chs[i] = make(chan struct{})
	}
	for i := 0; i < length; i++ {
		go func(index int, sender, receiver chan struct{}) {
			for {
				token := <-sender
				time.Sleep(3 * time.Second)
				fmt.Println(index)
				receiver <- token
			}
		}(i+1, chs[i], chs[(i+1)%length])
	}
	chs[0] <- struct{}{}
	select {}
}
