package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan interface{})
	//close done after 10 second
	time.AfterFunc(10*time.Second, func() { close(done) })

	const timeout = 2 * time.Second
	heartbeat, results := dowork(done, timeout/2)
	for {
		select {
		case _, ok := <-heartbeat:
			if !ok {
				return
			}
			fmt.Println("pulse")
		case r, ok := <-results:
			if !ok{
				return
			}
			fmt.Printf("resullt %v",r.Second())
		case <-time.After(timeout):
			fmt.Println("WORKER GOROUTINEIS not HESALHTY")
			return
		}
	}

}
