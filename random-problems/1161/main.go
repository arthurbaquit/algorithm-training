package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxLevelSum(root *TreeNode) int {
	levelSums := []int{}
	dfs(root, 0, &levelSums)

	maxSum := levelSums[0]
	idx := 1
	for i := 1; i < len(levelSums); i++ {
		if levelSums[i] > maxSum {
			maxSum = levelSums[i]
			idx = i + 1
		}
	}

	return idx
}

func dfs(node *TreeNode, level int, levelSums *[]int) {
	if node == nil {
		return
	}

	if len(*levelSums) == level {
		(*levelSums) = append(*levelSums, node.Val)
	} else {
		(*levelSums)[level] += node.Val
	}

	dfs(node.Left, level+1, levelSums)
	dfs(node.Right, level+1, levelSums)
}

func main() {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 7}
	root.Right = &TreeNode{Val: 0}
	root.Left.Left = &TreeNode{Val: 7}
	root.Left.Right = &TreeNode{Val: -8}
	println(maxLevelSum(root))
}
