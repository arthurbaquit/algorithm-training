package main

func sumOddLengthSubarrays(arr []int) int {
	l := len(arr)
	total := 0
	
	for i := 0; i < l; i++ {
		left, right := i,l - i - 1;
		total += arr[i] * (left / 2 + 1) * (right / 2 + 1);
		total += arr[i] * ((left + 1) / 2) * ((right + 1) / 2);
	}
	
	return total;
}

func main(){
	println(sumOddLengthSubarrays([]int{1,4,2,5,3}))
	println(sumOddLengthSubarrays([]int{1,2}))
	println(sumOddLengthSubarrays([]int{10,11,12}))
}