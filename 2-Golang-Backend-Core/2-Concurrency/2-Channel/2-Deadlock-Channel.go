package main

import (
	"fmt"
	"os"
	"time"
)

func sender_bad(c chan int) {
	for i := 1; i <= 3; i++ {
		c <- i
		time.Sleep(time.Millisecond * 300)
	}
}

func sender_good(c chan int) {
	for i := 1; i <= 3; i++ {
		c <- i
		time.Sleep(time.Millisecond * 300)
	}
	close(c)
}

func deadlock_channel() {
	fmt.Println("=== RUNNING: DEADLOCK CHANNEL ===")
	c := make(chan int)
	go sender_bad(c)

	for {
		fmt.Println("Received:", <-c)
	}
}

func no_deadlock_channel() {
	fmt.Println("=== RUNNING: NO DEADLOCK CHANNEL ===")
	c := make(chan int)
	go sender_good(c)

	for i := range c {
		fmt.Println("Received:", i)	
	}
	fmt.Println("Channel was Closed!")
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("[ERROR]: Missing Arg")
		return
	}

	mode := os.Args[1]

	switch mode {
	case "deadlock":
		deadlock_channel()
	case "no_deadlock":
		no_deadlock_channel()
	default:
		fmt.Println("[ERROR]: 'deadlock' or 'no_deadlock'")
	}	
}
