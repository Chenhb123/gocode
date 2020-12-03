package main

import (
	"fmt"
	"time"
)

func main() {

	for {
		fmt.Printf("start %v \n", time.Now())
		time.Sleep(2 * time.Second)
		// time.Sleep的过程中，下面的打印不会执行
		fmt.Printf("end %v \n", time.Now())
	}
}
