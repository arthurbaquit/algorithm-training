package main

/**
 * Definition for singly-linked list.
 */

type ListNode struct {
    Val int
    Next *ListNode
}

func getDecimalValue(head *ListNode) int {
	decimal := 0
	for head != nil {
		decimal = 2*decimal + head.Val
		head = head.Next
	}
	return decimal
}

func main(){
	println(getDecimalValue(&ListNode{1, &ListNode{0, &ListNode{1, nil}}}))
}