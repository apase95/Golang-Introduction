package main

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"os/exec"
	"time"
)

func main() {
	fmt.Println("=== 1. BASIC COMMAND EXECUTION ===")
	cmd1 := exec.Command("echo", "Hello World!")
	out1, err := cmd1.Output()
	if err != nil {
		log.Fatalf("Command execution failed: %v", err)
	}; fmt.Printf("[Basic_Output]: %s", string(out1))


	fmt.Println("\n=== 2. CAPTURING STDOUT AND STDERR ===")
	cmd2 := exec.Command("ls", "/non_existent_directory")
	var stdout, stderr bytes.Buffer
	cmd2.Stdout = &stdout
	cmd2.Stderr = &stderr
	err = cmd2.Run()
	if err != nil {
		fmt.Printf("[Command_Failed] Expected error caught: %v\n", err)
		fmt.Printf("[STDERR_Details]: %s", stderr.String())
	} else {
		fmt.Printf("[STDOUT_Details]: %s\n", stdout.String())
	}


	fmt.Println("\n=== 3. ADVANCED: COMMAND WITH TIMEOUT ===")
	ctx, cancel := context.WithTimeout(context.Background(), 6 * time.Second)
	defer cancel()
	cmd3 := exec.CommandContext(ctx, "sleep", "5")
	fmt.Println("Running a 5-second task with a 2-second timeout...")
	err = cmd3.Run()
	if ctx.Err() == context.DeadlineExceeded {
		fmt.Println("[TIMEOUT_ALERT] Command took too long and was forcefully killed!")
	} else if err != nil {
		fmt.Printf("Other error occurred: %v\n", err)
	} else {
		fmt.Println("Command finished successfully on time.")
	}
}