package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
简单的channel、协程并发
*/

func main() {
	channel := make(chan string)
	go producer("dog", channel)
	go producer("cat", channel)
	customer(channel)
}

// 生产者
func producer(header string, channel chan<- string) {
	for {
		channel <- fmt.Sprintf("%s:%d", header, rand.Int31())
		time.Sleep(1 * time.Second)
	}
}

// 消费者
func customer(channel <-chan string) {
	for {
		message := <-channel
		fmt.Println(message)
	}
}
