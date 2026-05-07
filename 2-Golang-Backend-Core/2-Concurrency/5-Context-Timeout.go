package main

import (
	"context"
	"fmt"
	"time"
)

func executeLongTask(ctx context.Context, taskName string) error {
	fmt.Printf("[TASK] Starting '%s' (requires 3 seconds to complete)...\n", taskName)
	select {
	case <- time.After(3 * time.Second):
		fmt.Printf("[TASK] '%s' completed successfully!\n", taskName)
		return nil
	
	case <- ctx.Done():
		fmt.Printf("[TASK_ERROR] '%s' was aborted: %v\n", taskName, ctx.Err())
		return ctx.Err()
	}
}

func main() {
	fmt.Println("=== SCENARIO 1: TASK FAILS DUE TO TIMEOUT ===")
	ctx1, cancel1 := context.WithTimeout(context.Background(), 2 * time.Second)
	defer cancel1()
	_ = executeLongTask(ctx1, "Download Database Dump")

	fmt.Println("\n=== SCENARIO 2: TASK COMPLETE BEFORE TIMEOUT ===")
	ctx2, cancel2 := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel2()
	_ = executeLongTask(ctx2, "Run Unit Tests")
}