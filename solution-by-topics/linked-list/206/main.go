package main
 type ListNode struct {
      Val int
      Next *ListNode
  }
 func reverseList(head *ListNode) *ListNode {
    var prev *ListNode
    curr:= head
    for curr != nil {
        next := curr.Next
        curr.Next = prev
        prev = curr
        curr = next
    }
    return prev
}

func main(){
	res := reverseList(&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}})
	for res!= nil{
		print(res.Val, " ")
		res = res.Next
	}
	println()
}