package main

import "fmt"

func floodFill(image [][]int, sr int, sc int, color int) [][]int {
    if image[sr][sc] == color { return image }
    row, col := len(image), len(image[sr])
    actualColor := image[sr][sc]
    image[sr][sc] = color
    if sr+1 < row && image[sr+1][sc] == actualColor { floodFill(image, sr+1, sc, color)}    
    if sr-1 >= 0 && image[sr-1][sc] == actualColor { floodFill(image, sr-1, sc, color)}
    if sc+1 < col && image[sr][sc+1] == actualColor { floodFill(image, sr, sc+1, color)}
    if sc-1 >= 0 && image[sr][sc-1] == actualColor { floodFill(image, sr, sc-1, color)}
    return image    
}

func main(){
	fmt.Println(floodFill([][]int{{1,1,1},{1,1,0},{1,0,1}}, 1, 1, 2))
}