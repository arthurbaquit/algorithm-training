package main

import "fmt"

/**
 * Definition for a binary tree node.
 */
 type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

 func postorderTraversal(root *TreeNode) []int {
    res := []int{}
    if root == nil {return res}
    res = append(res, postorderTraversal(root.Left)...)
    res = append(res, postorderTraversal(root.Right)...)
    res = append(res, root.Val)
    return res
}

func main(){
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 5}
	root.Right.Left = &TreeNode{Val: 6}
	root.Right.Right = &TreeNode{Val: 7}
	fmt.Println(postorderTraversal(root))
}