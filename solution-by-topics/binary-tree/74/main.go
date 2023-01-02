package main
func searchMatrix(matrix [][]int, target int) bool {
    row, col := len(matrix), len(matrix[0])
    ci, ce := 0, row*col-1
    for ci <= ce {
        mid := ci + (ce-ci)/2
        midElement := matrix[mid / col][mid % col];
        if midElement == target { return true}
        if midElement < target { ci = mid+1; continue}
        ce = mid-1
    }
    return false
}

func main(){
	println(searchMatrix([][]int{{1,3,5,7},{10,11,16,20},{23,30,34,60}},3))
	println(searchMatrix([][]int{{1,3,5,7},{10,11,16,20},{23,30,34,60}},13))
}