package main

import "fmt"

func matrixReshape(mat [][]int, r int, c int) [][]int {
    rows := len(mat)
    col := len(mat[0])
    if rows*col != r*c { return mat}
    res := make([][]int, r)
    for i:=0;i<r;i++{
        res[i] = make([]int, c)
    }
    r1, r2 := -1, -1
    for i:=0;i<r*c;i++{
        c1, c2 := i%c, i%col
        if (c1 == 0) {r1++}
        if (c2 == 0) {r2++}
        res[r1][c1] = mat[r2][c2]
    }
    return res
}

func main(){
	fmt.Println(matrixReshape([][]int{{1,2},{3,4}},1,4))
	fmt.Println(matrixReshape([][]int{{1,2},{3,4}},2,4))
	fmt.Println(matrixReshape([][]int{{1,2},{3,4}},4,1))
	fmt.Println(matrixReshape([][]int{{1,2},{3,4}},2,3))
}