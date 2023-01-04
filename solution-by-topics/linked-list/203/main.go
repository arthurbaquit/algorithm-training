package main
type ListNode struct {
     Val int
     Next *ListNode
 }

 func removeElementsWithDummy(head *ListNode, val int) *ListNode {
    dummy := &ListNode{-1, head}
    prev, curr := dummy, head
    for curr != nil {
        if curr.Val == val {
            prev.Next = curr.Next
            curr = curr.Next
            continue
        }
        prev = curr
        curr = curr.Next
    }
    return dummy.Next
}

 func removeElementsWithoutDummy(head *ListNode, val int) *ListNode {
    for head != nil && head.Val==val {head = head.Next}
    aux := head
    for aux != nil && aux.Next != nil{
        if aux.Next.Val == val {
            aux.Next = aux.Next.Next
            continue
        }
        aux = aux.Next
    }
    return head
}

func main(){
	res := removeElementsWithDummy(&ListNode{1, &ListNode{2, &ListNode{6, &ListNode{3, &ListNode{4, &ListNode{5, &ListNode{6, nil}}}}}}}, 6)
	for res!= nil{
		print(res.Val, " ")
		res = res.Next
	}
	println()
	res = removeElementsWithoutDummy(&ListNode{1, &ListNode{2, &ListNode{6, &ListNode{3, &ListNode{4, &ListNode{5, &ListNode{6, nil}}}}}}}, 6)
	for res!= nil{
		print(res.Val, " ")
		res = res.Next
	}
}