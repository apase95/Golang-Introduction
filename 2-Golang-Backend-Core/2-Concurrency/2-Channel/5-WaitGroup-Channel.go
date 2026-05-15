package main

import (
	"fmt"
	"math/rand"
	"os"
	"sync"
	"time"
)

func no_waitgroup_channel() {
	go func() {
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		fmt.Println("Goroutine 1 done")
	}()
	go func() {
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		fmt.Println("Goroutine 2 done")
	}()
	time.Sleep(500 * time.Millisecond)
	// time.Sleep(1000 * time.Millisecond)
	// if time sleep < total(random time) -> only 1 on 2 (G1, G2) run
}

func waitgroup_channel() {
	wc := new(sync.WaitGroup)
	wc.Add(2)

	go func() {
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		fmt.Println("Goroutine 1 done")
		wc.Done()
	}()
	go func() {
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		fmt.Println("Goroutine 2 done")
		wc.Done()
	}()

	wc.Wait()
	fmt.Println("All Goroutines done")
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("[ERROR]: Missing Arg")
		return
	}

	mode := os.Args[1]

	switch mode {
	case "no_waitgroup":
		no_waitgroup_channel()
	case "waitgroup":
		waitgroup_channel()
	default:
		fmt.Println("[ERROR]: 'waitgroup' or 'no_waitgroup'")
	}
}