package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func delNodes(root *TreeNode, to_delete []int) []*TreeNode {
	to_del := make(map[int]bool)
	for _, node := range to_delete {
		to_del[node] = true
	}

	var result []*TreeNode

	if _, is_present := to_del[root.Val]; !is_present {
		result = append(result, root)
	}

	visit_postorder(root, &to_del, &result)
	return result
}

func visit_postorder(root *TreeNode, to_del *map[int]bool, result *[]*TreeNode) {
	if root == nil {
		return
	}

	visit_postorder(root.Right, to_del, result)
	visit_postorder(root.Left, to_del, result)

	if root.Left != nil {
		if _, is_present := (*to_del)[root.Left.Val]; is_present {
			root.Left = nil
		}
	}
	if root.Right != nil {
		if _, is_present := (*to_del)[root.Right.Val]; is_present {
			root.Right = nil
		}
	}

	if _, is_present := (*to_del)[root.Val]; is_present {
		if root.Left != nil {
			*result = append(*result, root.Left)
			root.Left = nil
		}
		if root.Right != nil {
			*result = append(*result, root.Right)
			root.Right = nil
		}
	}
}
func main() {
	root := &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 5}}, Right: &TreeNode{Val: 3, Left: &TreeNode{Val: 6}, Right: &TreeNode{Val: 7}}}
	result := delNodes(root, []int{3, 5})
	for _, node := range result {
		println(node.Val)
		for node.Left != nil {
			println(node.Left.Val)
			node = node.Left
		}
		for node.Right != nil {
			println(node.Right.Val)
			node = node.Right
		}

	}

}
