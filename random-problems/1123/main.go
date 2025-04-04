package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lcaDeepestLeaves(root *TreeNode) *TreeNode {
	ans, _ := dfs(root, 0)
	return ans
}

func dfs(root *TreeNode, currDepth int) (*TreeNode, int) {
	if root == nil {
		return nil, currDepth
	}
	maximumL, leftD := dfs(root.Left, currDepth+1)
	maximumR, rightD := dfs(root.Right, currDepth+1)
	if leftD == rightD {
		return root, leftD
	}
	if leftD > rightD {
		return maximumL, leftD
	}
	return maximumR, rightD
}

func main() {
	root := &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 5}}, Right: &TreeNode{Val: 3, Left: &TreeNode{Val: 6}, Right: &TreeNode{Val: 7}}}
	_ = lcaDeepestLeaves(root)
}
