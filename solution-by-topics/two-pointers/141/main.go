package main

/**
 * Definition for singly-linked list.
 */

type ListNode struct {
    Val int
    Next *ListNode
}

func hasCycle(head *ListNode) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {return true}
	}
	return false
}

func main(){
	cycle := &ListNode{3, &ListNode{2, &ListNode{0, &ListNode{-4, nil}}}}
	cycle.Next.Next.Next.Next = cycle.Next
	println(hasCycle(cycle))
	println(hasCycle(&ListNode{3, &ListNode{2, &ListNode{0, &ListNode{-4, nil}}}}))
	println(hasCycle(&ListNode{1, nil}))
}