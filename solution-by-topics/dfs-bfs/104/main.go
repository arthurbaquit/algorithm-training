package main

/**
 * Definition for a binary tree node.
 */

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

 func maxDepth(root *TreeNode) int {
    return dfs(root, 0)   
}

func dfs(root *TreeNode, count int) int{
    if root == nil {return count}
    count++
    c1, c2 := dfs(root.Left, count), dfs(root.Right, count)
    return Max(c1,c2)
}

func Max(x,y int) int {
    if x>y { return x}
    return y
}

func main(){
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 5}
	root.Right.Left = &TreeNode{Val: 6}
	root.Right.Right = &TreeNode{Val: 7}
	println(maxDepth(root))
}