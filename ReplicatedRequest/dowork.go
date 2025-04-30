package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var dowork = func(
	done <-chan interface{},
	id int,
	wg *sync.WaitGroup,
	result chan<- int,
) {
	started := time.Now()
	defer wg.Done()

	LoadTime := time.Duration(1+rand.Intn(5)) * time.Second

	select {
	case <-done:
	case <-time.After(LoadTime):
	}

	select {
	case <-done:
	case result <- id:
	}

	took := time.Since(started)

	if took < LoadTime {
		took = LoadTime
	}

	fmt.Printf("%v took %v\n", id, took)

}

func main() {
	done := make(chan interface{})
	result := make(chan int)

	var wg sync.WaitGroup
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go dowork(done, i, &wg, result)
	}

	firstReturned := <- result
	close(done)
	wg.Wait()

	fmt.Printf("REcieved id %v", firstReturned)

}
