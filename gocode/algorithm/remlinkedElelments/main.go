package main

import "fmt"

func main() {
	head := &ListNode{}
	head.Val = 10
	head.Next = &ListNode{
		Val: 12,
	}
	head.Next.Next = &ListNode{
		Val: 14,
	}
	head.Next.Next.Next = &ListNode{
		Val: 12,
	}
	head = removeElements(head, 12)
	fmt.Println(head.Next.Val)
}

// ListNode .
type ListNode struct {
	Val  int
	Next *ListNode
}

// head: 1->2->6->3->4->5->6->6 val: 6
// p: 0->... 1->... 2->... 2->3->... 3->... 4->... 5->... 5->6
// q: 1->... 2->... 6->... 3->4->... 4->... 5->... 6->... 6->nil
// head: 1->2->3->4->5
func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}
	p := head
	q := p.Next
	for q != nil {
		if q.Val == val {
			p.Next = q.Next
		} else {
			p = q
		}
		q = p.Next
	}
	if head.Val == val {
		head = head.Next
	}
	return head
}
