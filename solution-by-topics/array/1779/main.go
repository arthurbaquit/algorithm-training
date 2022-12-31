package main

import "fmt"

func nearestValidPoint(x int, y int, points [][]int) int {
    min := 100000
    minCord := -1
    for i, cord := range(points) {
        if (x == cord[0] || y == cord[1]){
            aux := Min(min, ManhattanDistance(x, cord[0], y, cord[1]))
            if (min > aux){
                min = aux
                minCord = i
            } 
        }
    }
    return minCord
}

func Min(x,y int) int {
    if x>y {return y}
    return x
}

func ManhattanDistance (x1, x2, y1,y2 int) int{
    return Abs(x1 - x2) + Abs(y1-y2)
}
func Abs(x int) int {
    if x < 0 { return -x}
    return x
}

func main(){
	fmt.Println(nearestValidPoint(3,4, [][]int{{1,2},{3,1},{2,4},{2,3},{4,4}}))
	fmt.Println(nearestValidPoint(3,4, [][]int{{3,4}}))
	fmt.Println(nearestValidPoint(3,4, [][]int{{2,3}}))
}
