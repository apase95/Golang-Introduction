package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
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

	var n, q int
	fmt.Fscanln(reader, &n, &q)

	var arr = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &arr[i])
	}
	fmt.Fprintln(writer, arr)

	for i := 0; i < q; i++ {
		var target int
		fmt.Fscan(reader, &target)

		var idx = slices.Index(arr, target)
		if idx != -1 {
			fmt.Fprintln(writer, idx)
		} else {
			fmt.Fprintln(writer, "Not Found")
		}
	}
}
