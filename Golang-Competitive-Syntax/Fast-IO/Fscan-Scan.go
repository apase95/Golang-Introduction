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

	var n int
	fmt.Fscan(reader, &n) 
	
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &arr[i])
	}

	for {
		b, err := reader.ReadByte()
		if err != nil || (b != ' ' && b != '\n' && b != '\r') {
			reader.UnreadByte()
			break
		}
	}

	fullString, _ := reader.ReadString('\n')	
	fullString = strings.TrimSpace(fullString)
	fmt.Fprintln(writer, fullString)
}