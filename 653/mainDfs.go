package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findTarget(root *TreeNode, k int) bool {
	numsMap := make(map[int]bool)
	return dfs(root, k, numsMap)
}

func dfs(root *TreeNode, k int, numsMap map[int]bool) bool {
	if root == nil {
		return false
	}
	if _, ok := numsMap[k-root.Val]; ok {
		return true
	}
	numsMap[root.Val] = true
	return dfs(root.Left, k, numsMap) || dfs(root.Right, k, numsMap)
}

func main() {
	root := &TreeNode{Val: 5}
	root.Left = &TreeNode{Val: 3}
	root.Right = &TreeNode{Val: 6}
	root.Left.Left = &TreeNode{Val: 2}
	root.Left.Right = &TreeNode{Val: 4}
	root.Right.Right = &TreeNode{Val: 7}
	println(findTarget(root, 9))
	println(findTarget(root, 28))
}
