package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)
	go func() {
		time.Sleep(1 * time.Second)
		fmt.Println("Goroutine done.")
		c <- 0
	}()
	fmt.Println("Waiting for goroutine...")

	fmt.Printf("Done! (%d)\n", <-c)
}
