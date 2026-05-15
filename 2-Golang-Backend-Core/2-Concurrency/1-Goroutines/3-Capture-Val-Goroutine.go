package main

import (
	"fmt"
	"os"
	"time"
)

func goroutine_capture(value int) {
	fmt.Println("=== RUNNING: CAPTURING VALUE IN GOROUTINE ===")
	for val := 1; val <= value; val++ {
		go func() {
			fmt.Println(val)
		}()
	}
	time.Sleep(time.Second)
}

func goroutine_no_capture(value int) {
	fmt.Println("=== RUNNING: NO CAPTURING VALUE IN GOROUTINE ===")
	for val := 1; val <= value; val++ {
		go func(i int) {
			fmt.Println(i)
		}(val)
	}
	time.Sleep(time.Second)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("[ERROR]: Missing Arg")
		return
	}

	mode := os.Args[1]

	switch mode {
	case "capture":
		goroutine_capture(20)
	case "no_capture":
		goroutine_no_capture(20)
	default:
		fmt.Println("[ERROR]: 'capture' or 'no_capture'")
	}
}

// go run 3-Capture-Val-Goroutine.go capture
// or
// go run 3-Capture-Val-Goroutine.go no_capture