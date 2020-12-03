package main

import (
	"fmt"
	"time"
)

func main() {
	for {
		timer := time.NewTimer(5 * time.Second)
		<-timer.C
		fmt.Println("Timer fired")
	}
}
