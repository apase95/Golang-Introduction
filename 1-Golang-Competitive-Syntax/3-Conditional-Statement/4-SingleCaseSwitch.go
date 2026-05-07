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

	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var q int
	fmt.Fscanln(reader, &q)
	for i := 0; i < q; i++ {
		var day int
		fmt.Fscanln(reader, &day)

		switch day {
			case 2:
				fmt.Fprintln(writer, "Monday")
			case 3:
				fmt.Fprintln(writer, "Tuesday")
			case 4:
				fmt.Fprintln(writer, "Wednesday")
			case 5:
				fmt.Fprintln(writer, "Thursday")
			case 6:
				fmt.Fprintln(writer, "Friday")
			case 7:
				fmt.Fprintln(writer, "Saturday")
			case 8:
				fmt.Fprintln(writer, "Sunday")
			default:
				fmt.Fprintln(writer, "Invalid day")
		}
	}
}