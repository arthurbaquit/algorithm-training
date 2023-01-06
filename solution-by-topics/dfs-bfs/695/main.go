package main

func maxAreaOfIsland(grid [][]int) int {
    max:= 0
    for i:=0;i<len(grid);i++{
        for j:=0;j<len(grid[i]);j++{
            count := dfs(grid, i,j, 0)
            max = Max(max, count)
        }
    }
    return max
}

func dfs (grid [][]int, x,y, countSize int) int {
    if grid[x][y] == 0 { return countSize }
    countSize++
    grid[x][y] = 0
    if x+1 < len(grid) { countSize = dfs(grid, x+1, y, countSize)}
    if y+1 < len(grid[x]) { countSize = dfs(grid, x, y+1, countSize)}
    if x-1 >= 0 { countSize = dfs(grid, x-1, y, countSize)}
    if y-1 >= 0 { countSize = dfs(grid, x, y-1, countSize)}
    return countSize
}

func Max(x,y int) int {
    if x>y{return x}
    return y
}

func main(){
	println(maxAreaOfIsland([][]int{{0,0,1,0,0,0,0,1,0,0,0,0,0},{0,0,0,0,0,0,0,1,1,1,0,0,0},{0,1,1,0,1,0,0,0,0,0,0,0,0},{0,1,0,0,1,1,0,0,1,0,1,0,0},{0,1,0,0,1,1,0,0,1,1,1,0,0},{0,0,0,0,0,0,0,0,0,0,1,0,0},{0,0,0,0,0,0,0,1,1,1,0,0,0},{0,0,0,0,0,0,0,1,1,0,0,0,0}}))
	println(maxAreaOfIsland([][]int{{0,0,0,0,0,0,0,0}}))
}