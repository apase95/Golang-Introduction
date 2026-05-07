package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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

	//reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var sb strings.Builder
	for i := 1; i <= 5; i++ {
		sb.WriteString("Line ")
		fmt.Fprintf(&sb, "%d\n", i)
	}

	res := sb.String()
	fmt.Fprintln(writer, res)
}