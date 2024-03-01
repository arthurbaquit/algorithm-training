package main

import "fmt"

func generatePossibleNextMoves(currentState string) []string {
	res := []string{}
	for i := 0; i < len(currentState)-1; i++ {
		if currentState[i] == '+' && currentState[i+1] == '+' {
			res = append(res, currentState[:i]+"--"+currentState[i+2:])
		}
	}
	return res
}

func main() {
	fmt.Println(generatePossibleNextMoves("++++"))
	fmt.Println(generatePossibleNextMoves("+++++"))
	fmt.Println(generatePossibleNextMoves("++++-"))
	fmt.Println(generatePossibleNextMoves("++-+"))
}
