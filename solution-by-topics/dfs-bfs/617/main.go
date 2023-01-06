package main

/**
 * Definition for a binary tree node.
 */

 type TreeNode struct {
     Val int
     Left *TreeNode
     Right *TreeNode
 }

 func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
    if root1 == nil { return root2 }
    if root2 == nil { return root1 }
    root := TreeNode{}
    root.Val = root1.Val + root2.Val
    root.Left = mergeTrees(root1.Left, root2.Left)
    root.Right = mergeTrees(root1.Right, root2.Right)
    return &root
}

func main(){
	res := mergeTrees(&TreeNode{1, &TreeNode{3, &TreeNode{5, nil, nil}, &TreeNode{2, nil, nil}}, &TreeNode{2, nil, nil}}, &TreeNode{2, &TreeNode{1, nil, &TreeNode{4, nil, nil}}, &TreeNode{3, nil, &TreeNode{7, nil, nil}}})
	println(res.Val)
	println(res.Left.Val)
	println(res.Right.Val)
	println(res.Left.Left.Val)
	println(res.Left.Right.Val)
	println(res.Right.Right.Val)
}
	