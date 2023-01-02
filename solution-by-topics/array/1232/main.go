package main

func checkStraightLine(coordinates [][]int) bool {
    x0, y0 := coordinates[0][0], coordinates[0][1]
    x1, y1 := coordinates[1][0], coordinates[1][1]
    for i:=2;i<len(coordinates);i++{
        currX, currY := coordinates[i][0], coordinates[i][1]
        if (x1 - x0) * (currY - y1) != (y1 - y0) * (currX - x1) { return false}
    }
    return true
}

func main(){
	println(checkStraightLine([][]int{{1,2},{2,3},{3,4},{4,5},{5,6},{6,7}}))
	println(checkStraightLine([][]int{{1,1},{2,2},{3,4},{4,5},{5,6},{7,7}}))
}