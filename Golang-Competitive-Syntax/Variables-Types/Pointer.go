package main

import (
	"bufio"
	"fmt"
	"os"
)

func changeValue(x int) {
	x = 100
}

func changePointer(x *int) {
	*x = 100
}

func main() {
	if _, err := os.Stat("TEST.INP"); err == nil {
		inFile, _ := os.Open("TEST.INP")
		defer inFile.Close()
		os.Stdin = inFile

		outFile, _ := os.Create("TEST.OUT")
		defer outFile.Close()
		os.Stdout = outFile
	}

	//reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	val := 10
	
	changeValue(val)
	fmt.Printf("After changeValue: %d\n", val)

	changePointer(&val)
	fmt.Printf("After changePointer: %d\n", val)
}