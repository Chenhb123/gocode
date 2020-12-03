package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	length := len(os.Args)
	if length != 2 {
		log.Fatal("错误的参数:", os.Args)
	}
	tag := os.Args[1]
	list := []int{1, 2, 3, 4, 5}
	var result int
	// Label:
	for _, v := range list {
		fmt.Printf("%d\t", v)
		switch tag {
		case "$":
			result += v
		case "#":
			result += 2 * v
		default:
			break
			// break Label

		}
	}
	fmt.Println(result)
}
