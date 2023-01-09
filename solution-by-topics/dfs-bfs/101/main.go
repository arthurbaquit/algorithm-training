package main

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
    return isMirror(root, root);
}

func isMirror(node1, node2 *TreeNode) bool {
    if node1 == nil && node2 == nil {return true}
    if node1 == nil || node2 == nil {return false}
    return (node1.Val == node2.Val) && isMirror(node1.Right, node2.Left) && isMirror(node1.Left, node2.Right)
}

func main(){
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 2}
	root.Left.Left = &TreeNode{Val: 3}
	root.Left.Right = &TreeNode{Val: 4}
	root.Right.Left = &TreeNode{Val: 4}
	root.Right.Right = &TreeNode{Val: 3}
	println(isSymmetric(root))
}