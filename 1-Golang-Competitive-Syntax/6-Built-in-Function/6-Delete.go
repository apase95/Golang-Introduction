package main

import (
	"bufio"
	"fmt"
	"os"
)

type Student struct {
	Name  string
	Score int
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

	var n int
	fmt.Fscan(reader, &n)

	var scores = make(map[string]int)
	var records []Student
	for i := 0; i < n; i++ {
		var name string
		var score int
		fmt.Fscan(reader, &name, &score)
		scores[name] = score
		records = append(records, Student{Name: name, Score: score})
	}

	delete(scores, "A")
	delete(scores, "C")	

	for name, score := range scores {
		fmt.Fprintf(writer, "%s: %d\n", name, score)
	}

	fmt.Fprintln(writer, "----------------------------------------")
	
	if len(records) > 0 {
		records = records[1:]
		fmt.Fprintln(writer, records)
	}
}