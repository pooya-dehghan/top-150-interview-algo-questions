package main

import "fmt"

// first step generator
func generator(list []int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, v := range list {
			out <- v
		}
	}()

	return out
}

// second step multiplier

func multiplier(mlCh <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		for v := range mlCh {
			out <- v * 2
		}
	}()

	return out
}

// third step printer

func printer(printCh <-chan int) {

	for v := range printCh {
		fmt.Printf(" %v : ", v)
	}
}

func main() {

	list := []int{1, 2, 3, 4, 5, 6, 7, 8}

	mlCh := generator(list)
	prCh := multiplier(mlCh)
	printer(prCh)
}
