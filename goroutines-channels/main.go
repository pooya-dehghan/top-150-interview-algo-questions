package main

import (
	"fmt"
	"sync"
)

func halfer(nums []int) ([]int, []int) {

	sliceLength := len(nums)

	firstHalve := nums[:sliceLength/2]

	secondHalve := nums[(sliceLength / 2):sliceLength]

	return firstHalve, secondHalve

}

func summation(nums []int, ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	var sums int

	for i := range len(nums) {
		sums += nums[i]
	}

	ch <- sums
}

func summationPrinter(ch chan int) {
	var totalSums int
	for value := range ch {
		totalSums += value
	}

	fmt.Printf("totalSum: %v", totalSums)
}

func main() {
	ch := make(chan int)
	wg := sync.WaitGroup{}
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8}

	firstHalv, secondHalv := halfer(nums)

	wg.Add(2)

	go summation(firstHalv, ch, &wg)
	go summation(secondHalv, ch, &wg)

	go func() {
		wg.Wait()
		close(ch)
	}()

	summationPrinter(ch)

}
