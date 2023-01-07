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

 func inorderTraversal(root *TreeNode) []int {
    res := []int{}
    if root == nil { return res }
    res = append(res, inorderTraversal(root.Left)...)
    res = append(res, root.Val)
    res = append(res, inorderTraversal(root.Right)...)
    return res
}

func inorderTraversalUsingMorrisTransversal(root *TreeNode) []int {
	res := []int{}
	if root == nil { return res }
	var current *TreeNode = root
	for current != nil {
		if current.Left == nil {
			res = append(res, current.Val)
			current = current.Right
			continue
		}

		var pre *TreeNode = current.Left
		for pre.Right != nil && pre.Right != current {
			pre = pre.Right
		}
		pre.Right, current, current.Left = current, current.Left, nil
		
	}
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
	fmt.Println(inorderTraversal(root))
	fmt.Println(inorderTraversalUsingMorrisTransversal(root))
}