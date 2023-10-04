package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	var prev *ListNode
	curr := slow.Next
	for curr != nil {
		aux := curr.Next
		curr.Next = prev
		prev = curr
		curr = aux
	}
	firstHalf := head
	secondHalf := prev
	for secondHalf != nil {
		if firstHalf.Val != secondHalf.Val {
			return false
		}
		firstHalf = firstHalf.Next
		secondHalf = secondHalf.Next
	}
	return true
}

func main() {
	head := &ListNode{Val: 1, Next: &ListNode{Val: 2}}
	fmt.Println(isPalindrome(head))
	head = &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 2}}}
	fmt.Println(isPalindrome(head))
	head = &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 1}}}
	fmt.Println(isPalindrome(head))
}
