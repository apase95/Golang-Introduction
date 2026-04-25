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

	var q int
	fmt.Fscan(reader, &q)
	for i := 0; i < q; i++ {
		var n int
		fmt.Fscan(reader, &n)
		if n%2 == 0 {
			fmt.Fprintln(writer, "EVEN")
		} else {
			fmt.Fprintln(writer, "ODD")
		}
	}
}