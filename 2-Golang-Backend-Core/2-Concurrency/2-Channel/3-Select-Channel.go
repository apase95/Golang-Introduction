package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func create_channel() (chan int, chan int) {
	ch1 := make(chan int)
	ch2 := make(chan int)
	
	go func() {
		delay := time.Duration(rand.Intn(1000)) * time.Millisecond
		time.Sleep(delay)
		ch1 <- 1
	}()

	go func() {
		delay := time.Duration(rand.Intn(1000)) * time.Millisecond
		time.Sleep(delay)
		ch2 <- 2
	}()
	return ch1, ch2
}

func no_select_channel() {
	ch1, ch2 := create_channel()

	fmt.Println(<-ch1)
	fmt.Println(<-ch2)
}

func select_channel() {
	ch1, ch2 := create_channel()

	select {
	case v1 := <-ch1:
		fmt.Println("Ch1 come first with value:", v1)
		fmt.Println("then ch2 with value:", <-ch2)
	case v2 := <-ch2:
		fmt.Println("Ch2 come first with value:", v2)
		fmt.Println("then ch1 with value:", <-ch1)
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("[ERROR]: Missing Arg")
		return
	}

	mode := os.Args[1]

	switch mode {
	case "no_select":
		no_select_channel()
	case "select":
		select_channel()
	default:
		fmt.Println("[ERROR]: 'select' or 'no_select'")
	}
}