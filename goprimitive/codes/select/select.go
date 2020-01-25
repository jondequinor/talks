package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	realizationsChan := make(chan int)
	for i := 1; i <= 100; i++ {
		go func(index int) {
			time.Sleep(time.Duration(rand.Intn(3000)) * time.Millisecond)

			realizationsChan <- index
		}(i)
	}

	for i := 1; i <= 100; i++ {
		fmt.Printf("r%d\n", <-realizationsChan)
	}
	fmt.Println("\nDone!")
}
