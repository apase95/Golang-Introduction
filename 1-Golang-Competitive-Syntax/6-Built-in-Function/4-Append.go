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

	var n, m int 
	fmt.Fscan(reader, &n, &m)

	var arr []int
	for i := 0; i < n; i++ {
		var x int
		fmt.Fscan(reader, &x)
		arr = append(arr, x)
	}

	var arr2 []int
	for i := 0; i < m; i++ {
		var x int
		fmt.Fscan(reader, &x)
		arr2 = append(arr2, x)
	}

	arr = append(arr, arr2...)
	arr = append(arr, 100, 200, 300)

	for _, x := range arr {
		fmt.Fprintln(writer, x)
	}
}