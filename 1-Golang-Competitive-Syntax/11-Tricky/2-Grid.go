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

	var n, m int
	fmt.Fscan(reader, &n, &m)


	//----------------GRID--------------------
	var grid = make([][]int, n)				//vector<vector<int>> grid(n, vector<int>(m, 0));
	for i := 0; i < n; i++ {				//Grid (n * m). Value = 0
		grid[i] = make([]int, m)
	}

	for i := 0; i < n; i++ {			
		for j := 0; j < m; j++ {
			fmt.Fscan(reader, &grid[i][j])  //Fill value into grid
		}
	}
}