package main

import (
	"fmt"
	"os"
	"time"
)

func sayHello(name string) {
	for i := 0; i < 5; i++ {
		fmt.Printf("Hello %s\n", name)
	}
}

func bad_gorountine(name1 string, name2 string) {
	fmt.Println("=== RUNNING: BAD GOROUTINES ===")
	go sayHello(name1)
	sayHello(name2)
}

func normal_goroutine(name1 string, name2 string) {
	fmt.Println("=== RUNNING: GOOD GOROUTINES ===")
	go sayHello(name1)
	sayHello(name2)
	time.Sleep(time.Second)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("[ERROR]: Missing Arg")
		return
	}

	mode := os.Args[1]

	switch mode {
	case "bad":
		bad_gorountine("VIET", "NAM")
	case "normal":
		normal_goroutine("VIET", "NAM")
	default:
		fmt.Println("[ERROR]: 'normal' or 'bad'")
	}
}

// go run 1-Normal-Goroutine.go normal
// or
// go run 1-Normal-Goroutine.go bad