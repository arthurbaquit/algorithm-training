package main


type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

 func insertIntoBST(root *TreeNode, val int) *TreeNode {
    if root == nil { return &TreeNode{Val:val, Right: nil, Left: nil}}
    curr := root
    for curr != nil{
        if curr.Val > val {
            if curr.Left == nil { curr.Left = &TreeNode{Val: val, Left: nil, Right: nil}; break }
            curr = curr.Left
            continue
        }
        if curr.Right == nil { curr.Right = &TreeNode{Val: val, Left: nil, Right: nil}; break }
           curr = curr.Right
    }
    return root
}

func main(){
	root := &TreeNode{Val: 4}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 7}
	root.Left.Left = &TreeNode{Val: 1}
	root.Left.Right = &TreeNode{Val: 3}
	root = insertIntoBST(root, 5)
	println(root.Right.Left.Val)
}