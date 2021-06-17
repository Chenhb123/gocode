package main

import "time"

func main() {
	go func() {
		println("hello world")
	}()

	time.Sleep(10 * time.Minute)
}
