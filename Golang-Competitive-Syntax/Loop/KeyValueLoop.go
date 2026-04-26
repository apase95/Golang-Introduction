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

	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscanln(reader, &n)
	
	a := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}

	for idx, val := range a {
		fmt.Fprintf(writer, "Index: %d, Value: %s\n", idx, val)
	}
}