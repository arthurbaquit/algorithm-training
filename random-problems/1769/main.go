package main

import "fmt"

func abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}

func minOperationsBruteForce(boxes string) []int {
	answer := make([]int, len(boxes))
	for i := range boxes {
		for j, b := range boxes {
			if b == '1' {
				answer[i] += abs(i - j)
			}
		}
	}
	return answer
}

func minOperations(boxes string) []int {
	answer := make([]int, len(boxes))
	leftBoxes := 0
	rightBoxes := 0
	cost := 0
	// as we are starting from idx 0
	// all boxes will be to its right
	// so the box counting will be all to rightBoxes
	// and the total cost is the index of that box
	for i, b := range boxes {
		if b == '1' {
			rightBoxes++
			cost += i
		}
	}
	for i, b := range boxes {
		answer[i] = cost
		if b == '1' {
			// if this box has 1, that means that, from now on,
			// this box will be at left from the current idx.
			// So the cost is the "left" cost instead of "right" cost
			leftBoxes++
			rightBoxes--
		}
		// as we walk one idx to the right, all left boxes costs one more,
		// and the right boxes costs one less. So we just increase the
		// cost based in the number of each leftBox and decrease for each rightBox
		cost += leftBoxes - rightBoxes
	}
	return answer
}

func main() {
	fmt.Println(minOperations("110"))              // [1,1,3]
	fmt.Println(minOperationsBruteForce("110"))    // [1,1,3]
	fmt.Println(minOperations("001011"))           // [11,8,5,4,3,4]
	fmt.Println(minOperationsBruteForce("001011")) // [11,8,5,4,3,4]
}
