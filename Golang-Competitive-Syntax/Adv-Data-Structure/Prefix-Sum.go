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

	var n, q int
	fmt.Fscan(reader, &n, &q)

	arr := make([]int, n + 1)
	pref := make([]int, n + 1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(reader, &arr[i])
		pref[i] = pref[i - 1] + arr[i]
	}
	fmt.Fprintln(writer, arr)
	fmt.Fprintln(writer, pref)

	for i := 0; i < q; i++ {
		var l, r int
		fmt.Fscan(reader, &l, &r)
		fmt.Fprintf(writer, "[%d -> %d] = %d\n", l, r, pref[r] - pref[l - 1])
	}
}