package main

import (
	"fmt"
	"os"
	"runtime"
	"time"
)

func goroutine_schedule(value int) {
	fmt.Println("=== RUNNING: GOROUTINE HAS SCHEDULE ===")
	go func() {
		for i := 1; i <= value; i++ {
			fmt.Println("G1")
		}
	}()

	go func() {
		for i := 1; i <= value; i++ {
			fmt.Println("G2")
		}
	}()
	time.Sleep(time.Second)
}

func goroutine_force_schedule(value int) {
	fmt.Println("=== RUNNING: GOROUTINE HASN'T SCHEDULE ===")
	go func() {
		for i := 1; i <= value; i++ {
			fmt.Println("G1")
			runtime.Gosched()
		}
	}()
	go func() {
		for i := 1; i <= value; i++ {
			fmt.Println("G2")
			runtime.Gosched()
		}
	}()
	time.Sleep(time.Second)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("[ERROR]: Missing Arg")
		return
	}

	mode := os.Args[1]

	switch mode {
	case "schedule":
		goroutine_schedule(5)
	case "force_schedule":
		goroutine_force_schedule(5)
	default:
		fmt.Println("[ERROR]: 'schedule' or 'force_schedule'")
	}
}

// go run 4-Goroutine-Gosched.go schedule
// or
// go run 4-Goroutine-Gosched.go force_schedule