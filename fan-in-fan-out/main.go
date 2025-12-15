package main

import "fmt"

func multiplier(ch <-chan int, sqCh chan<- int) {
	for val := range ch {
		sqCh <- val * val
	}

}

func main() {
	ch := make(chan int, 101)
	sqCh := make(chan int, 101)

	go multiplier(ch, sqCh)
	go multiplier(ch, sqCh)
	go multiplier(ch, sqCh)
	go multiplier(ch, sqCh)
	go multiplier(ch, sqCh)

	go func() {
		rep := 10
		fmt.Println("start")
		for i := 0; i < rep; i++ {
			ch <- i
			fmt.Printf(" %v", i)
		}
	}()

	for {
	}
}
