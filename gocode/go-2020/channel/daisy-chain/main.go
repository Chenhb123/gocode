package main

import (
	"fmt"
)

func f(left, right chan int) {
	left <- 1 + <-right
}

/*

right9->right8->right7->right6->right5->right4->right3->right2->right1->right0->leftmost
right0~9 实际上都是right，加上编号便于区分
现在赋值right <- 1(right9 <- 1)
1->2->3->4->5->6->7->8->9->10->leftmost:11
所以x值为11，此时所有channel都已完成读写操作，则再向right读写都会panic
*/
func main() {
	leftmost := make(chan int)
	var left, right chan int = nil, leftmost
	for i := 0; i < 10; i++ {
		left, right = right, make(chan int)
		go f(left, right)
	}
	right <- 1
	x := <-leftmost
	fmt.Println(x)
	//fmt.Println(<-right)  //为啥panic
	right <- 3 // panic
}
