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

	var intA, intB, intC int
	fmt.Fscan(reader, &intA, &intB, &intC)

	var strA, strB, strC string
	fmt.Fscan(reader, &strA, &strB, &strC)

	fmt.Fprintln(writer, min(intA, intB, intC))
	fmt.Fprintln(writer, min(strA, strB, strC))
}