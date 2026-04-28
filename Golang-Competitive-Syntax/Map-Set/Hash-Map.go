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
	fmt.Fscan(reader, &n)

	freq := make(map[int]int)
	for i := 0; i < n; i++ {
		var x int
		fmt.Fscan(reader, &x)
		freq[x]++
	}

	for key, value := range freq {
		fmt.Fprintf(writer, "[%d, %d]\n", key, value)
	}

	target := 30
	if count, exists := freq[target]; exists {
		fmt.Fprintf(writer, "%d -> %d\n", target, count)
	} else {
		fmt.Fprintf(writer, "%d not found in the map\n", target)
	}

	delete(freq, 20)
	fmt.Fprintln(writer, freq)
}