package main

import "fmt"

func multiplier(ch <-chan int, sqCh chan<- int) {

	select {
	case val := <-ch:
		sqCh <- val * val
		fmt.Println("output", val*val)
	}
}

func main() {
	ch := make(chan int)
	sqCh := make(chan int)

	go multiplier(ch, sqCh)
	go multiplier(ch, sqCh)
	go multiplier(ch, sqCh)
	go multiplier(ch, sqCh)
	go multiplier(ch, sqCh)

	i := 100

	for i < 100 {
		ch <- i

		i--
	}

}
