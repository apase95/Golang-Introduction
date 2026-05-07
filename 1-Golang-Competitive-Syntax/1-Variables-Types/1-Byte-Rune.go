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

	var charA byte = 'A'
	fmt.Fprintf(writer, "Char: %c has value %d\n", charA, charA)

	var charViet rune = 'Đ'
	fmt.Fprintf(writer, "Char: %c has value %d\n", charViet, charViet)

	s := "abc123"
	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			val := int(s[i] - '0')
			fmt.Fprintf(writer, "Digit: %d\n", val)
		} else {
			upper := s[i] - 32
			fmt.Fprintf(writer, "Uppercase: %c\n", upper)
		}
	}
}