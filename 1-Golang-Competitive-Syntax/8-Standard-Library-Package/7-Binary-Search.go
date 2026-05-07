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
	fmt.Fscan(reader, &n, &q)
	var arr = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &arr[i])
	}
	fmt.Fprintln(writer, arr)

	slices.Sort(arr)

	for i := 0; i < q; i++ {
		var target int
		fmt.Fscan(reader, &target)

		var idx, found = slices.BinarySearch(arr, target)
		if found {
			fmt.Fprintf(writer, "Target: %d -> Index: %d\n", target, idx)
		} else {
			fmt.Fprintf(writer, "Target: %d -> Not found\n", target)
		}
	}
}