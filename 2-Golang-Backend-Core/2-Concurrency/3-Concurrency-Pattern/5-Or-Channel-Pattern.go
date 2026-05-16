package main

import (
	"fmt"
	"time"
)

func Or(channels ...<-chan any) <-chan any {
	switch len(channels) {
	case 0: 
		closeCh := make(chan any)
		close(closeCh)
		return closeCh
	case 1:
		return channels[0]
	default:
		orDone := make(chan any)
		go func() {
			defer close(orDone)

			select {
			case <-channels[0]:
			case <-channels[1]:
			case <-Or(append(channels[2:], orDone)...):
			}
		}()
		return orDone
	}
}

func Sig(after time.Duration) <-chan any {
	c := make(chan interface{})
	go func() {
		defer close(c)
		time.Sleep(after)
	}()
	return c
}

func main() {
	start := time.Now()
	<-Or(
		Sig(2 * time.Hour),
		Sig(5 * time.Minute),
		Sig(1 * time.Second),
		Sig(1 * time.Hour),
		Sig(1 * time.Minute),
	)
	fmt.Printf("done after %v", time.Since(start))
}