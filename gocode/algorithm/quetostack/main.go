package main

import "fmt"

func main() {
	s := Constructor()
	s.Push(1)
	s.Push(2)
	s.Push(3)
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Empty())
}

// MyStack 双队列，出栈时调整
/*
1.定义两个辅助队列q1、q2以及一个栈顶元素top_element
2.push:将元素加入q1
3.pop：
    3.1.将q1全部元素出队，除最后一个元素外，其余入队q2，然后删除q1中已入队q2的元素
    3.2.更新top_element
    3.3.调换q1和q2
4.top：返回top_element
5.empty:判断q1的长度
*/
type MyStack struct {
	Q1, Q2     []int
	TopElement int
}

/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{}
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	this.Q1 = append(this.Q1, x)
	this.TopElement = x
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	length := len(this.Q1)
	for i := 0; i < length-1; i++ {
		// 获取栈底元素
		e := this.Q1[0]
		// i = len(this.Q1) - 2时即为当前栈顶元素
		this.TopElement = e
		// 加入Q2
		this.Q2 = append(this.Q2, e)
		// 删除
		this.Q1 = this.Q1[1:]
	}
	// 获取出栈元素
	target := this.Q1[0]
	// 调换
	this.Q1 = this.Q2
	this.Q2 = make([]int, 0)
	return target
}

// Top Get the top element. */
func (this *MyStack) Top() int {
	return this.TopElement
}

// Empty Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return len(this.Q1) == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
