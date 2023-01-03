package main

import "strconv"

func isValidSudoku(board [][]byte) bool {
    rowsMap := [9][9]bool{}
    colsMap := [9][9]bool{}
    gridMap := [9][9]bool{}
    
    for row:=0; row<9; row++ {
        for col:=0; col<9; col++ {
            val, err := strconv.Atoi(string(board[row][col]))
            if err != nil {
                continue
            }
            val--
            gridIndex := col/3 + (row/3) * 3
            if rowsMap[row][val] || colsMap[col][val] || gridMap[gridIndex][val] {
                return false
            } 
            rowsMap[row][val] = true
            colsMap[col][val] = true    
            gridMap[gridIndex][val] = true
        }
    }
    return true
}

func main(){
	println(isValidSudoku([][]byte{
		{'5','3','.','.','7','.','.','.','.'},
		{'6','.','.','1','9','5','.','.','.'},
		{'.','9','8','.','.','.','.','6','.'},
		{'8','.','.','.','6','.','.','.','3'},
		{'4','.','.','8','.','3','.','.','1'},
		{'7','.','.','.','2','.','.','.','6'},
		{'.','6','.','.','.','.','2','8','.'},
		{'.','.','.','4','1','9','.','.','5'},
		{'.','.','.','.','8','.','.','7','9'},
	}))
	println(isValidSudoku([][]byte{
		{'8','3','.','.','7','.','.','.','.'},
		{'6','.','.','1','9','5','.','.','.'},
		{'.','9','8','.','.','.','.','6','.'},
		{'8','.','.','.','6','.','.','.','3'},
		{'4','.','.','8','.','3','.','.','1'},
		{'7','.','.','.','2','.','.','.','6'},
		{'.','6','.','.','.','.','2','8','.'},
		{'.','.','.','4','1','9','.','.','5'},
		{'.','.','.','.','8','.','.','7','9'},
	}))
	println(isValidSudoku([][]byte{
		{'.','.','.','.','5','.','.','1','.'},
		{'.','4','.','3','.','.','.','.','.'},
		{'.','.','.','.','.','3','.','.','1'},
		{'8','.','.','.','.','.','.','2','.'},
		{'.','.','2','.','7','.','.','.','.'},
		{'.','1','5','.','.','.','.','.','.'},
		{'.','.','.','.','.','2','.','.','.'},
		{'.','2','.','9','.','.','.','.','.'},
		{'.','.','4','.','.','.','.','.','.'},
	}))
}