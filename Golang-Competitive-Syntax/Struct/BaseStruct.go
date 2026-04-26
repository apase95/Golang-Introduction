package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Point struct {
	X int
	Y int
}

func comparePoints(p1, p2 Point) bool {
	if p1.X == p2.X {
		return p1.Y < p2.Y
	}
	return p1.X < p2.X
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
	fmt.Fscanln(reader, &n)
	markPoint := make([]Point, n)

	for i := 0; i < n; i++ {
		var p Point
		fmt.Fscanln(reader, &p.X, &p.Y)
		markPoint[i] = p
	}
	
	sort.Slice(markPoint, func(i, j int) bool {
		return comparePoints(markPoint[i], markPoint[j])
	})

	for _, p := range markPoint {
		fmt.Fprintln(writer, p.X, p.Y)
	}
}