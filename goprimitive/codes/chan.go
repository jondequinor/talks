ch := make(chan int) // ch is type 'chan int'

ch <- x // send statement

x = <-ch // receive expression in an assignment

<-ch // recieve statement, discarded
