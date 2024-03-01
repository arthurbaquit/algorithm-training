package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Result struct {
	Nodes [][]int
}

func findLeaves(root *TreeNode) [][]int {
	res := Result{
		Nodes: [][]int{},
	}
	_ = dfs(root, &res)
	return res.Nodes

}

func dfs(root *TreeNode, res *Result) int {
	if root == nil {
		return -1
	}
	if root.Left == nil && root.Right == nil {
		insertVal(res, 0, root.Val)
		return 0 + 1
	}
	max := 0
	if root.Left != nil {
		left := dfs(root.Left, res)
		if left > max {
			max = left
		}
	}
	if root.Right != nil {
		right := dfs(root.Right, res)
		if right > max {
			max = right
		}
	}
	insertVal(res, max, root.Val)
	return max + 1
}

func insertVal(result *Result, level int, value int) {
	var row []int
	if len(result.Nodes) <= level {
		row = make([]int, 0)
		result.Nodes = append(result.Nodes, row)
	} else {
		row = result.Nodes[level]
	}

	row = append(row, value)
	result.Nodes[level] = row
}

func main() {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 5,
			},
		},
		Right: &TreeNode{
			Val: 3,
		},
	}
	res := findLeaves(root)
	for _, v := range res {
		for _, v2 := range v {
			print(v2, " ")
		}
		println()
	}

}
