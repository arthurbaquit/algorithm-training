package main

/**
 * Definition for singly-linked list.
 */

type ListNode struct {
    Val int
    Next *ListNode
}
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    if list1 == nil { return list2}
    if list2 == nil { return list1}
    prev := &ListNode{-1, nil}
    head := prev
    for list1 != nil && list2 != nil {
        if list1.Val <= list2.Val{
            prev.Next = list1
            list1 = list1.Next
            prev = prev.Next
            continue
        }
        prev.Next = list2
        list2= list2.Next
        prev = prev.Next
    }
    if list1 == nil { 
        prev.Next = list2
        return head.Next
    }
    prev.Next = list1
    return head.Next
}

func main(){
    res := mergeTwoLists(&ListNode{1, &ListNode{2, &ListNode{4, nil}}}, &ListNode{1, &ListNode{3, &ListNode{4, nil}}})
    for res!= nil{
        print(res.Val, " ")
        res = res.Next
    }
}