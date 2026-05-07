package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func checkPrime(n int) bool {
	if n == 2 || n == 3 || n == 5 {
		return true
	}
	
	if n <= 1 || n % 2 == 0 || n % 3 == 0 {
		return false
	}

	limit := int(math.Sqrt(float64(n))) + 1
	for i := 6; i <= limit; i += 6 {
		if n % (i - 1) == 0 || n % (i + 1) == 0 {
			return false
		}
	}

	return true
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
	fmt.Fscan(reader, &q)

	for i := 0; i < q; i++ {
		var n int
		fmt.Fscan(reader, &n)
		if (checkPrime(n)) {
			fmt.Fprintf(writer, "%d -> YES\n", n)
		} else {
			fmt.Fprintf(writer, "%d -> NO\n", n)
		}
	}
}