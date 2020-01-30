package main

import (
	"fmt"
	"time"
)

func main() {
	helloChan := make(chan string)

	go func() {
		time.Sleep(time.Second)
		helloChan <- "World"
	}()

	fmt.Printf("Hello %s\n", <-helloChan)
}
