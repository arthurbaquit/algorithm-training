package main

func subsetXORSum(nums []int) int {
	subSetsLen := 1
	for range nums {
		subSetsLen *= 2
	}
	subSets := make([][]int, subSetsLen)
	currLen := 1
	for _, num := range nums {
		for i := range currLen {
			tmp := make([]int, len(subSets[i])+1)
			copy(tmp, subSets[i])
			tmp[len(tmp)-1] = num
			subSets[currLen+i] = tmp
		}
		currLen *= 2
	}
	ans := 0
	for _, subset := range subSets {
		tmp := 0
		for _, num := range subset {
			tmp ^= num
		}
		ans += tmp
	}
	return ans
}

func subsetXORSumOptimized(nums []int) int {
	subSetsLen := 1
	for range nums {
		subSetsLen *= 2
	}
	ans := 0
	for i := range subSetsLen {
		tmp := 0
		for j, num := range nums {
			// this creates a mask for the subsets
			// if the jth bit is set, we include the num in the subset
			// if the jth bit is not set, we don't include the num in the subset
			if (i>>j)&1 == 1 {
				tmp ^= num
			}
		}
		ans += tmp
	}
	return ans
}

func subsetXORSumBitManipulation(nums []int) int {
	ans := 0
	for _, num := range nums {
		ans |= num
	}
	return ans << (len(nums) - 1)
}

func main() {
	println(subsetXORSum([]int{1, 3}))                // 6
	println(subsetXORSumOptimized([]int{1, 3}))       // 6
	println(subsetXORSumBitManipulation([]int{1, 3})) // 6
}
