package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fileName := "new_file.txt"

	fmt.Println("=== 1. WRITE TO A NEW FILE ===")
	initialData := []byte("[INFO] Job started successfully.\n")
	err := os.WriteFile(fileName, initialData, 0644)
	if err != nil {
		log.Fatalf("Failed to write file: %v", err)
	}; fmt.Println("File created and written successfully.")


	fmt.Println("\n=== 2. APPEND TO AN EXISTING FILE ===")
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Failed to open file for appending: %v", err)
	}
	file.WriteString("[WARNING] Memory usage is getting high.\n")
	file.WriteString("[ERROR] Process terminated unexpectedly!\n")
	file.Close()
	fmt.Println("Logs appended successfully.")


	fmt.Println("\n=== 3. READ ENTIRE FILE INTO MEMORY ===")
	content, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}; fmt.Printf("File Content:\n%s", string(content))


	fmt.Println("\n=== 4. READ FILE LINE-BY-LINE (MEMORY EFFICIENT) ===")
	readFile, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}; defer readFile.Close()

	scanner := bufio.NewScanner(readFile)
	lineNumber := 1
	fmt.Println("Reading line by line:")
	for scanner.Scan() {
		fmt.Printf("  Line %d: %s\n", lineNumber, scanner.Text())
		lineNumber++
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("Error scanning file: %v", err)
	}
}