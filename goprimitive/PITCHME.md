## Go Concurrency Primitives

_(cc-by-sa 4.0)_

+++
### Why Go Concurrency Primitives
@css[](I really enjoyed working with concurrent Go)

@css[fragment](I want to share some of my enthusiasm)

@css[fragment](And possibly show you why I'm enthusiastic)

---

@snap[midpoint]
Background
@snapend

---

@snap[north-west]
### What is Go

@css[text-08](Minimalist language, language spec is short and readable)

@css[fragment text-08](Compiled, strong/static typing, garbage collected)

@css[fragment text-08](Strongly opinionated)

@css[fragment text-08](No generics, not very enterprise friendly)

@css[fragment text-08](Fast, powerful—simple, readable)

@css[fragment text-11](You get to run cool concurrent code @emoji[fragment em-sunglasses]) 
@snapend

@snap[east]

![fragment, rotate=60deg](https://i.imgur.com/GbBQDBA.png)
@snapend

+++

@code[golang](goprimitive/codes/hello/hello-world.go)

@[1](Declares what package this belongs to)
@[3](Familiar import statement)
@[5-7](As expected)

+++?color=#ffffdd

<iframe width="100%" height="315" src="http://localhost:8080/p/oXGayDtoLPh" frameborder="0" allowfullscreen></iframe>

---


### What is concurrency (in CS)

@css[](Composition of independently executing activities)

@css[fragment](Not parallelism, which is the simultaneous execution)

![fragment, filter=invert](https://msl-network.readthedocs.io/en/latest/_images/concurrency_vs_parallelism.png)

@css[fragment text-08](So concurrency is about structure, parallelism is about execution)

+++

### What is concurrency (in Go)

@css[text-07](Two styles:) @css[fragment text-07](shared-memory multi threading) @css[fragment text-07](and communicating sequential processes or _CSP_)

![fragment, filter=invert,height=300](https://ptolemy.berkeley.edu/publications/papers/99/HMAD/html/images/csp1.gif)

@css[fragment text-07](In CSP a program is a parallel composition of processes with no shared state)

@css[fragment text-07](The processes communicate and synchronize using channels)

+++

@snap[midpoint]
So how does Go do that?
@snapend

---

### Primitives

```go fragment
go // the ability to execute functions concurrently
````

```go fragment
chan // allows communication between concurrent processes
````

```go fragment
select // a multi-way concurrent control switch
````

+++

#### Goroutines

```go fragment
f() // Call `f()`; wait for it to return
```
```go fragment
go f() // Creates a new goroutine, no waiting
```
@css[fragment text-08](_Goroutines_ are independently executing functions. Think threads.)

@css[fragment text-08](They are really _cheap_, but not free)

@css[fragment text-08](Goroutines are really _coroutines_)

@css[fragment text-08](Meaning the scheduler in Go is a _cooperating scheduler_)

@css[fragment text-08](So scheduling decisions are made based on _user-space events_ i.e. `go`, `chan`, `select`, `io`, etc…)

+++

A concurrent Hello World program

@code[golang](goprimitive/codes/concurrent-hello/concurrent-hello.go)

@[9-12](Creates a goroutine that calls this _anonymous function_ created using a _function literal_)
@[14](Runs in the `main` goroutine)

+++?color=#ffffdd

@snap[north-west span-50]
<iframe width="750" height="400" src="http://localhost:8080/p/9lqafTHt4_x" frameborder="0" allowfullscreen></iframe> 
@snapend

@snap[north-east span-50 text-left]
@css[text-black text-09](what will this print)
@css[fragment text-black text-09](okay so that did not work @emoji[fragment em-crying_cat_face])
@css[fragment text-black text-09](when `main` returns, the program is terminated)
@css[fragment text-black text-09](how do we fix this?)
@snapend

@snap[south]
@css[fragment text-black](the `main` goroutine and the anonymous goroutine needs synchronization)
@css[fragment text-black](but how???)
@snapend

---

#### Channels

```go
ch := make(chan int) // ch is type 'chan int'
```

```go
ch <- x // send statement
```

```go
x = <-ch // receive expression in an assignment
```

```go
<-ch // recieve statement, discarded
```



+++

@snap[north]
#### Using channels
@snapend

@snap[west code-power span-50]
@code[golang zoom-08](goprimitive/codes/concurrent-hello-fixed/concurrent-hello-fixed.go)
@snapend

@snap[east span-40 text-left]
@[9](We make a `chan string` channel capable of communicating strings between goroutines)
@[13](The other goroutine sends `World`)
@[16](`main` will wait here until it will sync up with the anonymous function)
@[0-100]
@snapend

+++?color=#ffffdd

<iframe width="100%" height="500" src="http://localhost:8080/p/7YGyI0T1R1W" frameborder="0" allowfullscreen></iframe>

+++

#### Closing and ranging

@css[text-07](Channels can be `closed`)

```go
// closing a channel indicates to all receivers that nothing more
// will be sent
close(ch)

// receivers can test this like so. If not ok, the channel is closed.
_, ok := <-ch
```

@css[fragment text-07](It is also possible to loop over a channel until it is closed)

```go fragment
for x := range ch { // not the lack of receive expression
    // do something with x
}
// ch is now closed
```

+++

#### Channel axioms

@snap[midpoint span-100 text-center]
@ol
1. A send to a nil channel blocks forever
1. A receive from a nil channel blocks forever
1. A send to a closed channel panics
1. A receive from a closed channel returns the zero value immediately
@olend
@snapend

---

### Select

@css[fragment text-08](Let's you control your program's behaviour by looking at multiple channels at once)

@css[fragment text-08](The program asks "who is ready to communicate?")

@css[fragment text-08](It is like a switch statement, but for send/receive operations on channels)

@css[fragment text-08](Let's look at an example)

@code[golang fragment](goprimitive/codes/select.go)

@snap[south-east span-40 text-left]
@[0-100]
@[2](This case is selected if there's anything on the `ch` channel)
@[3]()
@[4](This case is selected if somebody is ready to receive on `helloChan`)
@[5]()
@[6](This case is selected if `ch` is empty and nobody is receiving on `helloChan`)
@[7]()
@[1-19](Note that if both `ch` is populated, and somebody is read for receiving on `helloChan`, it is random which is chosen)
@snapend

+++

#### Another example

@snap[west span-50]
```go zoom-08
package main

import (
	"fmt"
	"time"
)

func main() {

	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()

```
@snapend

@snap[east span-50]
```go zoom-08
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
    }()
    
    for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}
}
```
@snapend

+++?color=#ffffdd

<iframe width="100%" height="500" src="http://localhost:8080/p/SvOYelqbZeC" frameborder="0" allowfullscreen></iframe>

---

### Pipelines

@css[fragment text-07](Channels can be used to connect goroutines together)

@css[fragment text-07](The output of one, is the input of another)

![fragment](https://i.imgur.com/jAN13mS.png)

@css[fragment text-07](Excellent for data processing where you want to extend/change parts over time)

+++

#### Squarer program (1)
```go
package main

import "fmt"

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	// Counter
	go func() {
		for x := 0; x < 100; x++ {
			naturals <- x
		}
		close(naturals)
	}()
```

+++

#### Squarer program (2)

```go
	// Squarer
	go func() {
		for x := range naturals {
			squares <- x * x
		}
		close(squares)
	}()

	// Printer (in main goroutine)
	for x := range squares {
		fmt.Println(x)
	}
}
```

@css[fragment text-07](Is of course far too trivial to warrant its own goroutines)

@css[fragment text-07](But it can help illustrate how useful pipelines in go are)

+++?color=#ffffdd

<iframe width="100%" height="500" src="http://localhost:8080/p/KzQ-VLkr8T3" frameborder="0" allowfullscreen></iframe>

+++

#### Squarer with fan-out (1)

```go
func counter(out chan<- int, target int) {
	defer close(out)
	for x := 0; x < target; x++ {
		out <- x
	}
}
```

@snap[text-left]
@[1](Channels are first-class citizens)
@[1](And they can be directional)
@[0-100]
@snapend

+++

#### Squarer with fan-out (2)

```go
func squarer(in <-chan int, out chan<- int) {
	defer close(out)
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
}
```

@snap[text-left]
@[3](A `WaitGroup` can be used if we want to wait for multiple goroutines to finish)
@[6](Here, the square operation is now going to be executed independently)
@[12](Finally, when all square operations are done, `Wait` will return)
@[0-100]
@snapend

+++

#### Squarer with fan-out (3)

```go
func main() {
	naturals := make(chan int)
	squares := make(chan int)

	go counter(naturals, 1000000)
	go squarer(naturals, squares, false)

	for x := range squares {
		fmt.Println(x)
	}
}
```

+++?color=#ffffdd

<iframe width="100%" height="500" src="http://localhost:8080/p/GPtRzGGwf-J" frameborder="0" allowfullscreen></iframe>



---
@snap[north]
### Conclusion
@snapend

@snap[midpoint text-center span-100]
@ul[](false)
- Go supports a neat concurrency model, makes it easy to express complex operations dealing with concurrent problemt
- Fun to reason and talk about complex concurrent programming
- Arguably different/better than parallelism—and shared memory multithreading
@ulend
@snapend

---

@snap[midpoint]
## thanks
@snapend
