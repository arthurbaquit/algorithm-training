package main

type ListNode struct {
	Val int
	Next *ListNode
}

 func deleteDuplicates(head *ListNode) *ListNode {
    if head == nil { return head}
    curr, next := head, head
    for curr != nil {
        for next!=nil && next.Val==curr.Val{next = next.Next}
        curr.Next = next
        curr = next
    }
    return head
}

func main(){
	res := deleteDuplicates(&ListNode{1, &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{3, nil}}}}})
	for res!= nil{
		print(res.Val, " ")
		res = res.Next
	}
	println()
}