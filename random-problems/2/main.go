package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	res := &ListNode{}
	head := res
	aux := 0
	for l1 != nil && l2 != nil {
		res.Next = &ListNode{}
		res = res.Next
		res.Val = (l1.Val + l2.Val + aux) % 10
		aux = (l1.Val + l2.Val + aux) / 10
		l1, l2 = l1.Next, l2.Next
	}
	var reminding *ListNode
	if l1 != nil {
		reminding = l1
	} else {
		reminding = l2
	}
	for reminding != nil {
		res.Next = &ListNode{}
		res = res.Next
		res.Val = (reminding.Val + aux) % 10
		aux = (reminding.Val + aux) / 10
		reminding = reminding.Next
	}
	if aux != 0 {
		res.Next = &ListNode{Val: 1, Next: nil}
	}
	return head.Next
}

func main() {
	l1 := &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3, Next: nil}}}
	l2 := &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4, Next: nil}}}
	tail := addTwoNumbers(l1, l2)
	for tail != nil {
		fmt.Println(tail.Val)
		tail = tail.Next
	}
}
