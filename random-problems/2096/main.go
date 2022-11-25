package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteMiddle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	var pre, slow, fast *ListNode
	pre, slow, fast = nil, head, head
	for fast != nil && fast.Next != nil {
		pre = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	pre.Next = pre.Next.Next
	return head
}

func main() {
	head := &ListNode{1, nil}
	head.Next = &ListNode{2, nil}
	head.Next.Next = &ListNode{3, nil}
	head.Next.Next.Next = &ListNode{4, nil}
	head.Next.Next.Next.Next = &ListNode{5, nil}
	head.Next.Next.Next.Next.Next = &ListNode{6, nil}
	head.Next.Next.Next.Next.Next.Next = &ListNode{7, nil}
	head.Next.Next.Next.Next.Next.Next.Next = &ListNode{8, nil}
	head.Next.Next.Next.Next.Next.Next.Next.Next = &ListNode{9, nil}
	head.Next.Next.Next.Next.Next.Next.Next.Next.Next = &ListNode{10, nil}
	res := deleteMiddle(head)
	for res != nil {
		fmt.Printf("%d ", res.Val)
		res = res.Next
	}
}
