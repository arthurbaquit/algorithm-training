package main
func climbStairs(n int) int {
    countCache := make(map[int]int)
    return dfs(0, n, countCache)

}

func dfs(initial, target int, countCache map[int]int) int {
    if initial == target { return 1}
    if initial > target {return 0}
    count:=0
    v, ok := countCache[initial+1]
    if ok {
        count+= v
    } else {
        aux := dfs(initial+1, target, countCache)
        countCache[initial+1] = aux
        count += aux
    }
     v, ok = countCache[initial+2]
    if ok {
        count+= v
    } else {
        aux := dfs(initial+2, target, countCache)
        countCache[initial+2] = aux
        count += aux
    }
    return count    
}

func main(){
	println(climbStairs(5))
	println(climbStairs(45))
}