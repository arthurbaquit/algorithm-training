package main

func validateBinaryTreeNodes(n int, leftChild []int, rightChild []int) bool {
	root := findRoot(n, leftChild, rightChild)
	if root == -1 {
		return false
	}
	seen := make(map[int]bool)
	seen[root] = true
	stack := []int{root}
	for len(stack) > 0 {
		curr := stack[0]
		stack = stack[1:]
		for _, child := range []int{leftChild[curr], rightChild[curr]} {
			if child == -1 {
				continue
			}
			if _, ok := seen[child]; ok {
				return false
			}
			stack = append(stack, child)
			seen[child] = true
		}
	}
	return len(seen) == n
}

func findRoot(n int, left, right []int) int {
	seen := make(map[int]bool)
	for i := 0; i < n; i++ {
		seen[left[i]] = true
		seen[right[i]] = true
	}
	for i := 0; i < n; i++ {
		if _, ok := seen[i]; !ok {
			return i
		}
	}
	return -1
}

func main() {
	n := 4
	leftChild := []int{1, -1, 3, -1}
	rightChild := []int{2, -1, -1, -1}
	println(validateBinaryTreeNodes(n, leftChild, rightChild))
	rightChild = []int{2, 3, -1, -1}
	println(validateBinaryTreeNodes(n, leftChild, rightChild))
	n = 2
	leftChild = []int{1, 0}
	rightChild = []int{-1, -1}
	println(validateBinaryTreeNodes(n, leftChild, rightChild))
	n = 6
	leftChild = []int{1, -1, -1, 4, -1, -1}
	rightChild = []int{2, -1, -1, 5, -1, -1}
	println(validateBinaryTreeNodes(n, leftChild, rightChild))
}
