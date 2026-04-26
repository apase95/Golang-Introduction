package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	if _, err := os.Stat("TEST.INP"); err == nil {
		inFile, _ := os.Open("TEST.INP")
		defer inFile.Close()
		os.Stdin = inFile

		outFile, _ := os.Create("TEST.OUT")
		defer outFile.Close()
		os.Stdout = outFile
	}

	// reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	s := make([]int, 5, 10)
	fmt.Fprintln(writer, "Slice: ", s)
	fmt.Fprintln(writer, "Length: ", len(s))
	fmt.Fprintln(writer, "Capacity: ", cap(s))

	f := make(map[string]int)
	fmt.Fprintln(writer, "Length: ", len(f))
}