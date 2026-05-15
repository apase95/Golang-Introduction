package main

import (
	"fmt"
	"os"
)

func deadlock_buffered_channel() {
	bufferedChan := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		bufferedChan <- i
	}
	fmt.Printf("BufferedChan has len = %d, cap = %d", len(bufferedChan), cap(bufferedChan))
	bufferedChan <- 6 //Deadlock
}

func order_of_buffered_channel() {
	bufferedChan := make(chan int, 5)
	for i := 1; i <=5; i++ {
		bufferedChan <- i
	}
	for i := 1; i <=5; i++ {
		fmt.Println(<-bufferedChan)
	}
}

func compare_buffered_unbuffered() {
	bufferedChan := make(chan int, 1)
	unbufferedChan := make(chan int)

	bufferedChan <- 1    // OK
	unbufferedChan <- 1	 // deadlock
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("[ERROR]: Missing Arg")
		return
	}

	mode := os.Args[1]

	switch mode {
	case "deadlock":
		deadlock_buffered_channel()
	case "order":
		order_of_buffered_channel()
	case "compare":
		compare_buffered_unbuffered()
	default:
		fmt.Println("[ERROR]: 'deadlock' or 'order' or 'compare")
	}		
}