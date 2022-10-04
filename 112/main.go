package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumNode(node *TreeNode, targetSum, currentSum int) bool {
	currentSum += node.Val
	if node.Left == nil && node.Right == nil {
		return targetSum == currentSum
	}
	left, right := false, false
	if node.Left != nil {
		left = sumNode(node.Left, targetSum, currentSum)
	}
	if node.Right != nil {
		right = sumNode(node.Right, targetSum, currentSum)
	}
	return left || right
}
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	return sumNode(root, targetSum, 0)
}

func main() {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	println(hasPathSum(root, 5))
	println(hasPathSum(root, 3))
}
