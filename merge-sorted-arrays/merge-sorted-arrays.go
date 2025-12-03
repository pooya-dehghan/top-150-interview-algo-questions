package main

import (
	"fmt"
	"sort"
)

func merge(nums1 []int, m int, nums2 []int, n int) {
	nums1 = nums1[:m]
	for _, el := range nums2[:n] {
		nums1 = append(nums1, el)
	}

	// in case you need to remove duplicates

	// seenMap := make(map[int]bool)
	// newArr := []int{}

	// for _, v := range nums1 {
	// 	if !seenMap[v] {
	// 		seenMap[v] = true
	// 		newArr = append(newArr, v)
	// 	}
	// }

	sort.Ints(nums1)

	// nums1 = newArr

	fmt.Println(nums1)
}

func main() {
	merge([]int{1, 1, 2, 3, 0, 0}, 2, []int{1, 2, 3}, 3)

}
