package main

import "fmt"

/*
简单队列实现
*/

func main() {
	queue := []int{}
	// enqueue 入队
	queue = append(queue, 1)
	// front 队首元素
	fmt.Println(queue[0])
	// dequeue 出队
	queue = queue[1:]
	// isEmpty 非空
	fmt.Println(len(queue) <= 0)
}
