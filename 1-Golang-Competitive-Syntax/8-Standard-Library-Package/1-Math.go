package main

import (
	"bufio"
	"fmt"
	"math"
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

	fmt.Fprintf(writer, "Sqrt of 16: %v\n", math.Sqrt(16))
	fmt.Fprintf(writer, "2^3 = %v\n", math.Pow(2, 3))
	fmt.Fprintf(writer, "Absolute value of -5.5: %v\n", math.Abs(-5.5))

	n := -10
    absN := int(math.Abs(float64(n)))
    fmt.Fprintf(writer, "Absolute value of -10: %v\n", absN)

	fmt.Fprintf(writer, "Ceil of 3.14: %v\n", math.Ceil(3.14))
	fmt.Fprintf(writer, "Floor of 3.99: %v\n", math.Floor(3.99))

	fmt.Fprintf(writer, "Log2 of 8: %v\n", math.Log2(8))
	fmt.Fprintf(writer, "Log10 of 100: %v\n", math.Log10(100))
	
	fmt.Fprintf(writer, "Pi: %v\n", math.Pi)
	fmt.Fprintf(writer, "e: %v\n", math.E)
}