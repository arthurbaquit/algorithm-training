package main

func removeDuplicatesTwoPointers(nums []int) int {
	slow, fast, l, count := 0, 1, len(nums), 0
	for fast < l {
		for fast < l && nums[slow] == nums[fast] {
			fast++
		}
		nums[count] = nums[slow]
		slow = fast
		fast++
		count++
	}
	if slow >= l {
		return count
	}
	nums[count] = nums[slow]
	return count + 1
}

func removeDuplicates(nums []int) int {
	count := 1
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] != nums[i+1] {
			nums[count] = nums[i+1]
			count++
		}
	}
	return count
}

func main() {
	println("Duplicates with for loop")
	nums := []int{1, 1, 2}
	println(removeDuplicates(nums))
	nums = []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	println(removeDuplicates(nums))
	nums = []int{1, 1, 1, 2, 2, 3}
	println(removeDuplicates(nums))
	nums = []int{1, 1}
	println(removeDuplicates(nums))
	nums = []int{1, 1, 2}
	println("Duplicates with two pointers + while loop")
	println(removeDuplicatesTwoPointers(nums))
	nums = []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	println(removeDuplicatesTwoPointers(nums))
	nums = []int{1, 1, 1, 2, 2, 3}
	println(removeDuplicatesTwoPointers(nums))
	nums = []int{1, 1}
	println(removeDuplicatesTwoPointers(nums))
}
