package main

/**
 * Definition for singly-linked list.
 */

 type ListNode struct {
     Val int
     Next *ListNode
 }

 func removeNthFromEnd(head *ListNode, n int) *ListNode {
    dummy := ListNode{}
    dummy.Next= head
    slow, fast := &dummy, &dummy
    for i:=1; i<=n+1;i++{
            fast = fast.Next
        }
    for fast!= nil {
        fast = fast.Next
        slow = slow.Next
    }
    slow.Next = slow.Next.Next
    return dummy.Next
}

func main(){
	res := removeNthFromEnd(&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}, 2)
	for res!= nil{
		print(res.Val, " ")
		res = res.Next
	}
}
