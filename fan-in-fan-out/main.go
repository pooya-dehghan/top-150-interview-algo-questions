package main

import "sync"

func multiplier(ch <-chan int, sqCh chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for val := range ch {
		sqCh <- val * val
	}

}

func main() {
	ch := make(chan int)
	sqCh := make(chan int)
	var wg sync.WaitGroup

	groupCount := 10

	wg.Add(groupCount)
	for j := 0; j <= groupCount; j++ {
		go multiplier(ch, sqCh, &wg)
	}

	go func() {
		rep := 10
		for i := 0; i < rep; i++ {
			ch <- i
		}
		close(ch)
	}()

	go func() {
		wg.Wait()
		close(sqCh)
	}()

}
