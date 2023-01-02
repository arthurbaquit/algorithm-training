package main

/**
 * Definition for singly-linked list.
 */
 type ListNode struct {
     Val int
     Next *ListNode
 }
 
 func middleNode(head *ListNode) *ListNode {
    slow, fast := head, head
    for fast != nil && fast.Next != nil{
        slow = slow.Next
        fast = fast.Next.Next
    }
    return slow
}

func main(){
	res := middleNode(&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}})
	for res!= nil{
		println(res.Val)
		res = res.Next
	}
}
