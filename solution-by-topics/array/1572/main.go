package main

func diagonalSum(mat [][]int) int {
    totalSum := 0
    l := len(mat)
    for i:=0;i<l;i++{
        totalSum+=(mat[i][i] + mat[i][l-1-i])
    }
    if l%2==1{
    return totalSum-mat[l/2][l/2]

    }
    return totalSum
}

func main(){
	println(diagonalSum([][]int{{1,2,3},{4,5,6},{7,8,9}}))
	println(diagonalSum([][]int{{1,1,1,1},{1,1,1,1},{1,1,1,1},{1,1,1,1}}))
	println(diagonalSum([][]int{{5}}))
}
