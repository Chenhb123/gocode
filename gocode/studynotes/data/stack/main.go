package main

import (
	"errors"
	"fmt"
)

var stack []int

func main() {
	stack = make([]int, 0, 5)
	for i := 0; i < 6; i++ {
		fmt.Printf("%d: %v %v\n", i, push(i), stack)
	}
	for i := 0; i < 6; i++ {
		res, err := pop()
		fmt.Printf("%d: %v %v\n", res, err, stack)
	}
}

func push(x int) error {
	n := len(stack)
	if n == cap(stack) {
		return errors.New("stack is full")
	}
	stack = stack[:n+1]
	stack[n] = x
	return nil
}

func pop() (int, error) {
	n := len(stack)
	if n == 0 {
		return 0, errors.New("stack is empty")
	}
	res := stack[n-1]
	stack = stack[:n-1]
	return res, nil
}
