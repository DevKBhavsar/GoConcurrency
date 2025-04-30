package main

import "time"

var dowork = func(
	done <-chan interface{},
	pulseInterval time.Duration,
) (<-chan interface{}, <-chan time.Time) {
	heartbeat := make(chan interface{})
	result := make(chan time.Time)

	go func() {
		// //setup channel 
		// defer close(heartbeat)
		// defer close(result)

		//Everypulse Interval there will be something to read on this channel
		pulse := time.Tick(pulseInterval)
		//simulating work
		workGen := time.Tick(2 * pulseInterval)
	
		//default added because no-one might return to pulsebeat and they only care about result
		sendPulse := func() {
			select {
			case heartbeat <- struct{}{}:
			default:
			}
		}

		sendResult := func(r time.Time) {
			for {
				select {
				case <-done:
					return
				case <-pulse:
					sendPulse()
				case result <- r:
					return
				}

			}
		}

		for i:=0;i<2;i++ {
			select {
			case <-done:
				return
			case <-pulse:
				sendPulse()
			case r := <-workGen:
				sendResult(r)
			}
		}
	}()
	return heartbeat,result
}
