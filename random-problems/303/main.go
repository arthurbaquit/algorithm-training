package main

type NumArray struct {
    nums []int
}


func Constructor(nums []int) NumArray {
    sum := make( []int, len(nums) + 1);
    for  i := 0; i < len(nums); i++ {
        sum[i + 1] = sum[i] + nums[i];
    }
    return NumArray{nums: sum}    
}


func (arr *NumArray) SumRange(left int, right int) int {
   return arr.nums[right+1] - arr.nums[left]
}


/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */