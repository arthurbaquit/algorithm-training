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
 func preorderTraversal(root *TreeNode) []int {
	if root == nil { return []int{}}
   value := []int{}
   value = append(value, root.Val)
   value = append(value, preorderTraversal(root.Left)...)
   value = append(value, preorderTraversal(root.Right)...)
   return value

}


func preorderTraversalUsingMorrisTransversal(root *TreeNode) []int{
	res := []int{}
	if root == nil { return res}
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
		if pre.Right == nil {
			res = append(res, current.Val)
			pre.Right = current
			current = current.Left
			continue
		}
		pre.Right = nil
		current = current.Right
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
	fmt.Println(preorderTraversal(root))
	fmt.Println(preorderTraversalUsingMorrisTransversal(root))
}