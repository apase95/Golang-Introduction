package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
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
	var arr = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &arr[i])
	}
	fmt.Fprintf(writer, "Org: %v\n", arr)
	fmt.Fprintf(writer, "---------------------------------\n\n")

	var cloneArr1 = make([]int, n)
	copy(cloneArr1, arr)
	cloneArr1[0] = 100
	cloneArr1[len(cloneArr1) - 1] = 200
	fmt.Fprintf(writer, "Clone with 'copy()': %v\n", cloneArr1)
	fmt.Fprintf(writer, "Check Org after clone: %v\n\n", arr)

	var cloneArr2 = slices.Clone(arr)
	cloneArr2[0] = 300
	cloneArr2[len(cloneArr2) - 1] = 400
	fmt.Fprintf(writer, "Clone with 'slices.Clone()': %v\n", cloneArr2)
	fmt.Fprintf(writer, "Check Org after clone: %v\n\n", arr)

	var cloneArr3 = arr
	cloneArr3[0] = 500
	cloneArr3[len(cloneArr3) - 1] = 600
	fmt.Fprintf(writer, "Clone with assignment: %v\n", cloneArr3)
	fmt.Fprintf(writer, "Check Org after clone: %v\n", arr)
}