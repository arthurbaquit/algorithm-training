package main


type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

 func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
    mapP := find(root, p.Val, []*TreeNode{})
    mapQ :=find(root, q.Val, []*TreeNode{})
    point:= 0
    for point < len(mapP) && point < len(mapQ) && mapP[point] == mapQ[point] { point++ } 
    if point == len(mapP) { return mapP[len(mapP)-1]}
    if point == len(mapQ) { return mapQ[len(mapQ)-1]}
    return mapP[point-1]
}

func find(curr *TreeNode, target int, visited []*TreeNode) []*TreeNode{
    if curr == nil {return visited}
    visited = append(visited, curr)
    if curr.Val == target {return visited}
    if curr.Val > target {return find(curr.Left, target, visited); }
    return find(curr.Right, target, visited)
}

func lowestCommonAncestorSimple(root, p, q *TreeNode) *TreeNode {
    if root == nil {return root}
    if p.Val > root.Val && q.Val > root.Val { return lowestCommonAncestor(root.Right, p, q)}
    if p.Val < root.Val && q.Val < root.Val { return lowestCommonAncestor(root.Left, p, q)}
    return root
}

func main() {
	root := &TreeNode{Val: 6}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 8}
	root.Left.Left = &TreeNode{Val: 0}
	root.Left.Right = &TreeNode{Val: 4}
	root.Left.Right.Left = &TreeNode{Val: 3}
	root.Left.Right.Right = &TreeNode{Val: 5}
	root.Right.Left = &TreeNode{Val: 7}
	root.Right.Right = &TreeNode{Val: 9}
	println(lowestCommonAncestor(root, root.Left, root.Left.Right.Right).Val)
	println(lowestCommonAncestor(root, root.Left, root.Left.Right.Left).Val)
	println(lowestCommonAncestorSimple(root, root.Left, root.Left.Right.Right).Val)
	println(lowestCommonAncestorSimple(root, root.Left, root.Left.Right.Left).Val)

}