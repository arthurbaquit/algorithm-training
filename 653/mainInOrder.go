package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findTarget(root *TreeNode, k int) bool {
	numsList := []int{}
	inOrder(root, &numsList)
	i := 0
	j := len(numsList) - 1
	for i < j {
		if numsList[i]+numsList[j] == k {
			return true
		}
		if numsList[i]+numsList[j] < k {
			i++
		} else {
			j--
		}
	}
	return false
}

func inOrder(root *TreeNode, numsList *[]int) {
	if root == nil {
		return
	}
	// Stack everyone from the left side until return nil, that means, find a leaf
	inOrder(root.Left, numsList)
	// Then, save it in the list
	*numsList = append(*numsList, root.Val)
	// Then, go to the right side to grab all numbers greater than the current
	// one and smaller than the previous
	inOrder(root.Right, numsList)
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
