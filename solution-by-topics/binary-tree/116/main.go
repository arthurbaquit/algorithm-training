package main

import "fmt"

/**
 * Definition for a Node.
 */
 type Node struct {
     Val int
     Left *Node
     Right *Node
     Next *Node
 }


 func connect(root *Node) *Node {
	if root == nil { return nil}
    leftmost := root
    for leftmost.Left != nil {
        curr := leftmost
        for curr != nil{
			if curr.Next != nil {
				fmt.Println(curr.Next.Val)
			} else {
				fmt.Println("nil")
			}
            curr.Left.Next = curr.Right
            if curr.Next != nil { curr.Right.Next = curr.Next.Left}
            curr = curr.Next
        }
        leftmost = leftmost.Left
    }
	for leftmost != nil && leftmost.Next != nil {
		fmt.Println(leftmost.Next.Val)
		leftmost = leftmost.Next
	}

    return root
}

func main() {
	connect(&Node{1, &Node{2, &Node{4, nil, nil, nil}, &Node{5, nil, nil, nil}, nil}, &Node{3, &Node{6, nil, nil, nil}, &Node{7, nil, nil, nil}, nil}, nil})
}