package main

import (
	"fmt"
	"sync"
	"time"
)

func counter(out chan<- int, target int) {
	defer close(out)
	for x := 0; x < target; x++ {
		out <- x
	}
}

func squarer(in <-chan int, out chan<- int, fanout bool) {
	defer close(out)
	if fanout {
		wg := &sync.WaitGroup{}
		for x := range in {
			wg.Add(1)
			go func(x int) {
				time.Sleep(100 * time.Millisecond)
				out <- x * x
				wg.Done()
			}(x)
		}
		wg.Wait()
	} else {
		for x := range in {
			time.Sleep(100 * time.Millisecond)
			out <- x * x
		}
	}
}

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	go counter(naturals, 1000000)
	go squarer(naturals, squares, false)

	// Printer (in main goroutine)
	for x := range squares {
		fmt.Println(x)
	}
}
