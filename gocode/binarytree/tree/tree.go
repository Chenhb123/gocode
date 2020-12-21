package binarytree

import (
	"container/list"
	"math"
)

// BinaryTree 二叉树
type BinaryTree struct {
	Val         int
	Left, Right *BinaryTree
}

// New 新建二叉树节点
func New(val int) *BinaryTree {
	return &BinaryTree{Val: val}
}

// PreOrder 前序遍历 根节点-左节点-右节点
func (t *BinaryTree) PreOrder() []int {
	res := []int{}
	if t == nil {
		return res
	}
	node := t
	stack := list.New()
	for node != nil || stack.Len() != 0 {
		for node != nil {
			res = append(res, node.Val)
			stack.PushBack(node)
			node = node.Left
		}
		if stack.Len() != 0 {
			v := stack.Back()
			node = v.Value.(*BinaryTree)
			node = node.Right
			stack.Remove(v)
		}
	}
	return res
}

// MiddleOrder 中序遍历 左节点-根节点-右节点
func (t *BinaryTree) MiddleOrder() []int {
	res := []int{}
	if t == nil {
		return res
	}
	node := t
	stack := list.New()
	for node != nil || stack.Len() != 0 {
		for node != nil {
			stack.PushBack(node)
			node = node.Left
		}
		if stack.Len() != 0 {
			v := stack.Back()
			node = v.Value.(*BinaryTree)
			res = append(res, node.Val)
			node = node.Right
			stack.Remove(v)
		}
	}
	return res
}

// SufOrder 后序遍历 左节点-右节点-根节点
func (t *BinaryTree) SufOrder() []int {
	res := []int{}
	if t == nil {
		return res
	}
	node := t
	stack := list.New()
	var preNode *BinaryTree
	for node != nil || stack.Len() != 0 {
		for node != nil {
			stack.PushBack(node)
			node = node.Left
		}
		v := stack.Back()
		top := v.Value.(*BinaryTree)
		if (top.Left == nil && top.Right == nil) ||
			(top.Right == nil && preNode == top.Left) ||
			top.Right == preNode {
			res = append(res, top.Val)
			preNode = top
			stack.Remove(v)
		} else {
			node = top.Right
		}
	}
	return res
}

// LevelOrder 层次遍历
func (t *BinaryTree) LevelOrder() []int {
	res := []int{}
	if t == nil {
		return res
	}
	nodes := []*BinaryTree{t}
	var node *BinaryTree
	for len(nodes) != 0 {
		node = nodes[0]
		res = append(res, node.Val)
		nodes = nodes[1:]
		if node.Left != nil {
			nodes = append(nodes, node.Left)
		}
		if node.Right != nil {
			nodes = append(nodes, node.Right)
		}
	}
	return res
}

// IsBalanced 判断root是否为高度平衡二叉树
// 一个高度平衡二叉树的定义为：
// 一个二叉树每个节点的左右两个子树的高度差的绝对值不超过1
func IsBalanced(root *BinaryTree) bool {
	if root == nil {
		return true
	}
	lenfun := func(node *BinaryTree) int {
		var length int
		nodes := []*BinaryTree{}
		if node != nil {
			nodes = append(nodes, node)
			for len(nodes) != 0 {
				newnodes := []*BinaryTree{}
				length++
				for len(nodes) != 0 {
					newnode := nodes[0]
					nodes = nodes[1:]
					if newnode.Left != nil {
						newnodes = append(newnodes, newnode.Left)
					}
					if newnode.Right != nil {
						newnodes = append(newnodes, newnode.Right)
					}
				}
				nodes = newnodes
			}
		}
		return length
	}
	ll, rl := lenfun(root.Left), lenfun(root.Right)
	if math.Abs(float64(ll-rl)) > 1 {
		return false
	}
	return IsBalanced(root.Left) && IsBalanced(root.Right)
}

// SliLevelOrder  层次遍历，每层的结果为slice
func (t *BinaryTree) SliLevelOrder() [][]int {
	result := [][]int{}
	node := t
	nodes := []*BinaryTree{node}
	for len(nodes) != 0 {
		newnodes := []*BinaryTree{}
		res := []int{}
		for len(nodes) != 0 {
			tmpnode := nodes[0]
			nodes = nodes[1:]
			res = append(res, tmpnode.Val)
			if tmpnode.Left != nil {
				newnodes = append(newnodes, tmpnode.Left)
			}
			if tmpnode.Right != nil {
				newnodes = append(newnodes, tmpnode.Right)
			}
		}
		nodes = newnodes
		result = append(result, res)
	}
	return result
}

func MaxDepth(t *BinaryTree) int {
	if t == nil {
		return 0
	} else {
		left := MaxDepth(t.Left)
		right := MaxDepth(t.Right)
		if left > right {
			return 1 + left
		} else {
			return 1 + right
		}
	}
}
