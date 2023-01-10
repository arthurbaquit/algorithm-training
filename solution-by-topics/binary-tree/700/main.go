package main

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}
 func searchBST(root *TreeNode, val int) *TreeNode {
    if root == nil { return nil}
    if val == root.Val {return root}
    if root.Val > val { return searchBST(root.Left, val) }
    return searchBST(root.Right, val)
}

func main(){
	root := &TreeNode{Val: 4}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 7}
	root.Left.Left = &TreeNode{Val: 1}
	root.Left.Right = &TreeNode{Val: 3}
	head := searchBST(root, 2)
	if head != nil { println("Founded") } else { println("Not founded") }
}