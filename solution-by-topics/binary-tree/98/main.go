package main

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}
 func isValidBST(root *TreeNode) bool {
	return validade(root, nil, nil)
 }
 
 func validade (curr, high, low *TreeNode) bool {
	 if curr == nil { return true }
	 if (low != nil && curr.Val <= low.Val) || (high!=nil && curr.Val >= high.Val) { return false}
	 return validade(curr.Left, curr, low) && validade(curr.Right, high, curr)
 }

 func main(){
	root := &TreeNode{Val: 5}
	root.Left = &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 4}
	root.Right.Left = &TreeNode{Val: 3}
	root.Right.Right = &TreeNode{Val: 6}
	println(isValidBST(root))
 }