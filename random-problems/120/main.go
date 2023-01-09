package main

import (
	"fmt"
	"sort"
)

func minimumTotal(triangle [][]int) int {
    for row := 1; row < len(triangle); row++{
        for col := 0; col <len(triangle[row]); col++{
            smallestAbove := 100000
            if col > 0 { smallestAbove = triangle[row - 1][col - 1] }
            if col < len(triangle[row])-1 { smallestAbove = min(smallestAbove, triangle[row-1][col])} 
            triangle[row][col] = smallestAbove + triangle[row][col]
        }
    }
    sort.Ints(triangle[len(triangle)-1])
    return triangle[len(triangle)-1][0]
}
func min (x,y int)int {
    if x>y {return y}
    return x
}

func main(){
	fmt.Println(minimumTotal([][]int{{2},{3,4},{6,5,7},{4,1,8,3}}))
}