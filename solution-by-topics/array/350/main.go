package main

import "fmt"

func intersect(nums1 []int, nums2 []int) []int {
    mapNum := make(map[int]int)
    sol := make([]int, 0)
    for _, v := range(nums1) { 
        if _, ok := mapNum[v]; !ok {
            mapNum[v] = 1
            continue
        }
        mapNum[v]++
    }
     for _, v := range(nums2) { 
         if _, ok := mapNum[v]; ok && mapNum[v]>0 {
            sol = append(sol, v)
            mapNum[v]--
         }
    }

    return sol
}
func Min(x, y int) int {
    if x<y {return x}
    return y
}

func main(){
	nums1 := []int{1,2,2,1}
	nums2 := []int{2,2}
	fmt.Println(intersect(nums1, nums2))
	nums1 = []int{4,9,5}
	nums2 = []int{9,4,9,8,4}
	fmt.Println(intersect(nums1, nums2))
}

// Follow up:
// What if the given array is already sorted? How would you optimize your algorithm?
// I would use two pointers strategy. If nums1[i] == nums2[j] then add to solution and increment both pointers. If nums1[i] < nums2[j] then increment i, else increment j.
// That would be O(n+m) time complexity but it would be O(1) space complexity.

// What if nums1's size is small compared to nums2's size? Which algorithm is better?
// And the arrays are sorted? I would use the two pointers strategy. Because, on mean, it would break earlier than running through the whole nums2.
// Otherwise, of course I would stick with the map solution, as the two pointers strategy needs the sorted array.

// What if elements of nums2 are stored on disk, and the memory is limited such that you cannot load all elements into the memory at once?
// if nums1 fit the memory, using the map solution, I can interact as I'm loading portions of nums2 and check the map value.
// if nums1 doesn't fit the memory, I would partionate nums1, store an subarray into the map, then interact through nums2 (loading it dinamically). 
// Perfoms this until we checked all the subarrays of nums1.
