package main

import (
	"bufio"
	"fmt"
	"os"
)

func sum(a int, b int) (res int) {
	res = a + b
	return
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

	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var q int
	fmt.Fscanln(reader, &q)

	for i := 0; i < q; i++ {
		var a, b int
		fmt.Fscanln(reader, &a, &b)
		fmt.Fprintln(writer, sum(a, b))
	}
}