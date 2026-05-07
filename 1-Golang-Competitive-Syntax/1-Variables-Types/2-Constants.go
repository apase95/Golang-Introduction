package main

import (
	"bufio"
	"fmt"
	"os"
)

const MOD = 1e9 + 7
const INF = 1e18
const (
	MaxN = 1e5 + 5
	Base = 31
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
	fmt.Printf("a = %d, b = %f, a + MOD = %d\n", a, b, a + MOD)
	fmt.Printf("a = %d, b = %f, b + MOD = %d\n", a, b, b + MOD)
}