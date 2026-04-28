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

	//reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	s := "CodeForces"
	fmt.Fprintf(writer, "[Start to 4] %s\n", s[:4])
	fmt.Fprintf(writer, "[4 to End] %s\n", s[4:])
	fmt.Fprintf(writer, "[2 to 7] %s\n", s[2:7])

	bytesArr := []byte(s)
	bytesArr[0] = 'N'
	sNew := string(bytesArr)
	fmt.Fprintf(writer, "Modified String: %s\n", sNew)
}