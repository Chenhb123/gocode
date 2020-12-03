package main

import (
	"fmt"
	"time"
)

func main() {
	c := time.Tick(2 * time.Second)
	for next := range c {
		fmt.Printf("%v \n", next)
	}
}
