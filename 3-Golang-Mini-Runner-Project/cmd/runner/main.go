package main

import (
	"flag"
	"fmt"

	"golang-mini-runner-project/internal/config"
)

func main() {
	workersCount := flag.Int("workers", 5, "Number of concurrent workers in the pool")
	debugMode := flag.Bool("debug", false, "Enable debug mode logging")
	flag.Parse()
	fmt.Printf("[FLAG] Workers Count : %d\n", *workersCount)
	fmt.Printf("[FLAG] Debug Mode    : %v\n", *debugMode)
	
	appConfig := config.LoadConfig()
	config.PrintConfig(appConfig)
}