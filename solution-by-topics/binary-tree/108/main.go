package main
 type TreeNode struct {
     Val int
     Left *TreeNode
     Right *TreeNode
 }
 func sortedArrayToBST(nums []int) *TreeNode {
    if len(nums) == 0 { return nil }
    mid := len(nums)/2
    head := TreeNode{
        Val: nums[mid],
        Left: nil,
        Right: nil,
    }
    head.Left = sortedArrayToBST(nums[0:mid])
    head.Right = sortedArrayToBST(nums[mid+1:])
    return &head
}
func main() {
	head := sortedArrayToBST([]int{-10,-3,0,5,9})
	println(head.Val)
	println(head.Left.Val)
	println(head.Right.Val)
	println(head.Left.Left.Val)
	println(head.Right.Left.Val)
	//
}