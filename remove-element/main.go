package main

import "fmt"

func removeElement(nums []int, val int) int {
	arr := nums[:0]

	for _, el := range nums {
		if el != val {
			arr = append(arr, el)
		}
	}

	fmt.Println(arr)

	return len(arr)
}

func removeElementWaySecond(nums []int, val int) int {
	pointerOne := 0

	for pointerTwo := 0; pointerTwo < len(nums); pointerTwo++ {
		if nums[pointerTwo] != val {
			nums[pointerOne] = nums[pointerTwo]
			pointerOne++
		}
	}

	return pointerOne
}

func main() {
	removeElement([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2)
}
