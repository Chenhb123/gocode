package main

import (
	"fmt"
	"time"
)

func main() {

	for {
		dur := 2 * time.Second
		ticker := time.NewTicker(dur)
		done := make(chan bool)
		go func() {
			time.Sleep(5 * time.Second) // 5s后触发time.Stop，关闭ticker
			done <- true
		}()
		for {
			select {
			case <-done:
				fmt.Println("Done!")
				ticker.Stop()
				return
			case t := <-ticker.C:
				fmt.Println("Current time: ", t)
			}
		}
	}
}
