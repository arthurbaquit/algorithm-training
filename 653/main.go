package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findTarget(root *TreeNode, k int) bool {
	cache := make(map[int]bool)
	if root == nil {
		return false
	}
	return lookFor(root, root, k, cache)
}

func lookFor(root *TreeNode, actual *TreeNode, k int, cache map[int]bool) bool {
	if actual == nil {
		return false
	}
	if find(root, actual, k, cache) {
		return true
	}
	return lookFor(root, actual.Left, k, cache) || lookFor(root, actual.Right, k, cache)
}

func find(root *TreeNode, actual *TreeNode, k int, cache map[int]bool) bool {
	if root == nil {
		return false
	}

	if root == actual {
		if k-actual.Val-root.Val > 0 {
			return find(root.Right, actual, k, cache)
		}
		return find(root.Left, actual, k, cache)
	}
	if _, exists := cache[k-actual.Val]; !exists {
		if k-actual.Val == root.Val {
			return true
		}
		cache[root.Val] = !(k-root.Val == root.Val)
		if k-actual.Val-root.Val > 0 {
			return find(root.Right, actual, k, cache)
		}
		return find(root.Left, actual, k, cache)
	}
	return cache[k-actual.Val]
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
