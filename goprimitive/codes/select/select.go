package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	ch := make(chan int)
	helloChan := make(chan string)

	go func() {
		time.Sleep(time.Duration(rand.Intn(2)) * time.Second)
		fmt.Println(<-helloChan)
	}()

	go func() {
		time.Sleep(time.Duration(rand.Intn(2)) * time.Second)
		ch <- 42
	}()

	select {
	case x := <-ch:
		fmt.Printf("Got %d.", x)
	case helloChan <- "Hello World":
	}
}
