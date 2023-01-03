package main

func maximumWealth(accounts [][]int) int {
    max := 0
    for i:=0; i<len(accounts); i++{
        accountSum := 0
        for j:=0;j<len(accounts[i]);j++{
            accountSum+=accounts[i][j]
        }
        max = Max(accountSum,max)
    }
    return max
}

func Max(x,y int) int {
    if x>y {return x}
    return y
}

func main(){
	println(maximumWealth([][]int{{1,2,3},{3,2,1}}))
	println(maximumWealth([][]int{{1,5},{7,3},{3,5}}))
	println(maximumWealth([][]int{{2,8,7},{7,1,3},{1,9,5}}))
}