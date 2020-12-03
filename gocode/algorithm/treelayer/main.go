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

func (t *tree) levelOrder() [][]int {
	res := [][]int{}
	if t == nil {
		return res
	}
	nodes := []*tree{t}
	for len(nodes) != 0 {
		newnodes := []*tree{}
		layer := []int{}
		for _, v := range nodes {
			layer = append(layer, v.val)
			if v.left != nil {
				newnodes = append(newnodes, v.left)
			}
			if v.right != nil {
				newnodes = append(newnodes, v.right)
			}
		}
		res = append(res, layer)
		nodes = newnodes
	}
	return res
}

func create(value int) *tree {
	return &tree{val: value}
}

func main() {
	t := create(1)
	t.left, t.right = create(2), create(2)
	t.left.left, t.left.right = create(3), create(4)
	t.right.left, t.right.right = create(4), create(3)
	t.left.left.left, t.left.left.right = create(5), create(4)
	t.left.right.left, t.left.right.right = create(6), create(5)
	t.right.left.left, t.right.left.right = create(5), create(6)
	t.right.right.left, t.right.right.right = create(4), create(5)
	fmt.Println(t.preOrder())
	res := t.levelOrder()
	fmt.Println(res)
}
