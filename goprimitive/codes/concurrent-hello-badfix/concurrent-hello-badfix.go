package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		time.Sleep(time.Second)
		fmt.Print(" World\n")
	}()

	fmt.Print("Hello")
	time.Sleep(2 * time.Second)
}
