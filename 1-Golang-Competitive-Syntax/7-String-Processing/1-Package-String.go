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

	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	line1, _ := reader.ReadString('\n')
	line2, _ := reader.ReadString('\n')
	line1 = strings.TrimSpace(line1)
	line2 = strings.TrimSpace(line2)

	arr := strings.Split(line2, ",")
	fmt.Fprintln(writer, "Array after splitting:", len(arr), "elements")

	newStr := strings.Join(arr, " - ")
	fmt.Fprintln(writer, "String after joining:", newStr)

	if strings.Contains(line1, "Golang") {
		fmt.Fprintln(writer, "Line 1 contains 'Golang'")
	}

	count := strings.Count(line1, "o")
	fmt.Fprintln(writer, "Number of 'o' in Line 1:", count)

	separator := strings.Repeat("-", 20)
	fmt.Fprintln(writer, separator)
}