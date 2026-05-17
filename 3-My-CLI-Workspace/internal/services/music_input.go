package services

import (
	"bufio"
	"os"
	"strings"
)

func startKeyboardListener() <-chan string {
	controlChan := make(chan string)
	
	go func() {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			input := strings.TrimSpace(strings.ToLower(scanner.Text()))
			controlChan <- input
		}
	}()
	
	return controlChan
}