select {
case x := <-ch:
	fmt.Printf("Got %d.", x)
case helloChan <- "Hello World":
	fmt.Println("Sent Hello World.")
default:
	time.Sleep(time.Day)
}