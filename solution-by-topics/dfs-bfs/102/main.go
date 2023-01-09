package main

import "fmt"


type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}
func levelOrder(root *TreeNode) [][]int {
    var levels [][]int
    if (root == nil) { return levels}
    var dfs func(*TreeNode, int)
    dfs = func(node *TreeNode, level int ){
        if len(levels) == level { levels = append(levels, []int{})}
        levels[level] = append(levels[level], node.Val)
        if node.Left != nil {dfs(node.Left, level+1)}
        if node.Right != nil {dfs(node.Right, level+1)}
    }
    dfs(root, 0);
    return levels;
}

func main(){
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 5}
	root.Right.Left = &TreeNode{Val: 6}
	root.Right.Right = &TreeNode{Val: 7}
	fmt.Println(levelOrder(root))
}