package main

import "fmt"
func merge(nums1 []int, m int, nums2 []int, n int)  {
    if m == 0 { 
        for i:=0; i<m+n;i++{
        nums1[i] = nums2[i]
    }}
    if n == 0 {return}
    result := make([]int,0)
    pNum1, pNum2 := 0,0
    for pNum1 < m && pNum2 < n {
        for pNum1 < m && nums1[pNum1] <= nums2[pNum2]{
            result = append(result, nums1[pNum1])
            pNum1++
        }
        for pNum2 < n && nums2[pNum2] <= nums1[pNum1]{
            result = append(result, nums2[pNum2])
            pNum2++
        }
    }
    for i:= pNum1; i<m; i++{ result = append(result, nums1[i]) }
    for i:= pNum2; i<n; i++{ result = append(result, nums2[i]) }
    for i:=0; i<m+n;i++{
        nums1[i] = result[i]
    }

}

func main(){
	num1 := []int{1,2,3,0,0,0}
	merge(num1, 3, []int{2,5,6}, 3)
	fmt.Println(num1)
	num1 = []int{1}
	merge(num1, 1, []int{}, 0)
	fmt.Println(num1)
	num1 = []int{0}
	merge(num1, 0, []int{1}, 1)
	fmt.Println(num1)
	num1 = []int{4,5,6,0,0,0}
	merge(num1, 3, []int{1,2,3}, 3)
	fmt.Println(num1)
}