package main

import "fmt"

func generate(numRows int) [][]int {
    return recurse(generateBase(numRows))
}

func generateBase(rows int) [][]int {
    res := make([][]int, rows)
    
    for col := 0; col < rows; col++ {
        res[col] = make([]int, col+1)
        res[col][0], res[col][col] = 1, 1
    }

    res[0][0] = 1
    
    return res
}

func recurse(base [][]int) [][]int {
    rows := len(base)
    for row := 1; row < rows; row++ {
        for col := 1; col < len(base[row])-1; col++ {
            base[row][col] = base[row-1][col-1] + base[row-1][col]
        }
    }
    return base
}

func main(){
	fmt.Println(generate(5))
	fmt.Println(generate(1))
}