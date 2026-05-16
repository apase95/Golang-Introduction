package main

import (
	"fmt"
)

func generator(done <-chan struct{}, integers ...int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, i := range integers {
			select {
			case <-done:
				return
			case out <- i:
			}
		}
	}()
	return out
}

func add(done <-chan struct{}, in <-chan int, additive int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			select {
			case <-done:
				return
			case out <- n + additive:
			}
		}
	}()
	return out
}

func multiply(done <-chan struct{}, in <-chan int, multiplier int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			select {
			case <-done:
				return
			case out <- n * multiplier:
			}
		}
	}()
	return out
}

func main() {
done := make(chan struct{})
	defer close(done)
	
	intStream := generator(done, 1, 2, 3, 4)
	pipeline := multiply(done, add(done, multiply(done, intStream, 2), 1), 2)

	for v := range pipeline {
		fmt.Printf("%d ", v)
	}
}