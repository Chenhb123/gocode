package main

import "fmt"

type tree struct {
	val         int
	left, right *tree
}

func (t *tree) print() {
	fmt.Printf("%d \n", t.val)
}

func (t *tree) value() int {
	return t.val
}

func (t *tree) setValue(value int) {
	t.val = value
}

// 根节点-左子树-右子树
func (t *tree) preOrder() string {
	s := ""
	if t == nil {
		return ""
	}
	s = fmt.Sprintf("%s %d", s, t.val)
	s += t.left.preOrder()
	s += t.right.preOrder()
	return s
}

// 左子树-根节点-右子树
func (t *tree) middleOrder() string {
	s := ""
	if t == nil {
		return ""
	}
	s += t.left.middleOrder()
	s = fmt.Sprintf("%s %d", s, t.val)
	s += t.right.middleOrder()
	return s
}

func create(value int) *tree {
	return &tree{val: value}
}

/*

		1
	2		2
  3   4   4   3

*/

func main() {
	t := create(1)
	t.left, t.right = create(2), create(2)
	t.left.left, t.left.right = create(3), create(4)
	t.right.left, t.right.right = create(4), create(3)
	s := t.preOrder()
	fmt.Println(s)
}
