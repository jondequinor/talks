package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

type iteration struct {
	index int
	param int
	fm    []string
}

func generate(outCh chan<- iteration) {
	defer close(outCh)
	for i := 0; i < 100; i++ {
		outCh <- iteration{index: i, fm: []string{"copy_dir", "copy_dir", "copy_file"}}
	}
}

func modelparameters(inCh <-chan iteration, outCh chan<- iteration) {
	defer close(outCh)
	sampled := 0
	for iter := range inCh {
		iter.param = rand.Intn(1000)
		outCh <- iter
		sampled++
	}
	fmt.Printf("sampled %d iters\n", sampled)
}

func forwardmodel(inCh <-chan iteration, outCh chan<- iteration) {
	defer close(outCh)
	var ops uint64
	var wg sync.WaitGroup
	for iter := range inCh {
		wg.Add(1)
		go func(i iteration) {
			for range i.fm {
				time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
				atomic.AddUint64(&ops, 1)
			}
			wg.Done()
			outCh <- i
		}(iter)
	}
	wg.Wait()
	fmt.Printf("ran %d fms\n", ops)
}

func main() {
	sampleChan := make(chan iteration)
	forwardmodelChan := make(chan iteration)
	doneChan := make(chan iteration)

	go generate(sampleChan)
	go modelparameters(sampleChan, forwardmodelChan)
	go forwardmodel(forwardmodelChan, doneChan)

	for range doneChan {

	}
}
