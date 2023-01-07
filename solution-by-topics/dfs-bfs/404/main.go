package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}
func sumOfLeftLeaves(root *TreeNode) int {
    count := 0
    if root == nil {return count}
    if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil { 
        count+= root.Left.Val
        count+= sumOfLeftLeaves(root.Right)
        return count
    }
    count+=sumOfLeftLeaves(root.Left)
    count+=sumOfLeftLeaves(root.Right)
    return count
}
func main(){
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 5}
	root.Right.Left = &TreeNode{Val: 6}
	root.Right.Right = &TreeNode{Val: 7}
	println(sumOfLeftLeaves(root))
}