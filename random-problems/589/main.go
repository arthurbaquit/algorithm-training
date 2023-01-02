package main

import "fmt"

/**
 * Definition for a Node.
 */

  type Node struct {
      Val int
      Children []*Node
  }

 func preorder(root *Node) []int {
    res := []int{}
    if root == nil{
        return res
    }
	
    res = append(res,root.Val)
    for _, child := range root.Children{
         res = append(res, preorder(child)...)
    }
	
    return res
}

func main(){
	fmt.Println(preorder(&Node{1, []*Node{{3, []*Node{{5, nil},{6, nil}}},{2, nil}, {4, nil}}}))
}