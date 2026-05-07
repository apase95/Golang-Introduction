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

	var n int
	fmt.Fscan(reader, &n)

	var arr = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &arr[i])
	}

	//-----------QUEUE----------------
	var q []int  								//queue<int> q
	
	for i := 0; i < n; i++ {
		q = append(q, arr[i]) 					//st.push(arr[i])
	}
	fmt.Fprintln(writer, q)

	if len(q) > 0 {								//st.empty()
		fmt.Fprintln(writer, q[0]) 				//st.top()
	}
	
	q = q[1:] 									//st.pop()
	fmt.Fprintln(writer, q)
}
