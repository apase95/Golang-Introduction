package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

	var a int = 10
	var b float64 = 5.5
	c := float64(a) + b
	d := a + int(b)
	fmt.Fprintf(writer, "c = %f\n", c)
	fmt.Fprintf(writer, "d = %d\n", d)

	num := 123
	str := strconv.Itoa(num)
	fmt.Fprintf(writer, "String: %s (Type: %T)\n", str, str)

	newStr := "456"
	newNum, err := strconv.Atoi(newStr)
	if err == nil {
		fmt.Fprintf(writer, "Integer: %d (Type: %T)\n", newNum, newNum)
	}
}