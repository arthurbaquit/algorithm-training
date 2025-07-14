package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func getDecimalValue(head *ListNode) int {
	ans := 0
	for head != nil {
		ans *= 2
		ans += head.Val
		head = head.Next
	}
	return ans
}

func main() {
	head := &ListNode{Val: 1, Next: &ListNode{Val: 0, Next: &ListNode{Val: 1}}}
	println(getDecimalValue(head)) // 5

	head = &ListNode{Val: 0}
	println(getDecimalValue(head)) // 0

}
