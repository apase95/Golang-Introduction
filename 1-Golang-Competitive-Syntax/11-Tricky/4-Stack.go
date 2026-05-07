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

	//-----------STACK----------------
	var st []int  								//stack<int> st
	
	for i := 0; i < n; i++ {
		st = append(st, arr[i]) 				//st.push(arr[i])
	}
	fmt.Fprintln(writer, st)

	if len(st) > 0 {							//st.empty()
		fmt.Fprintln(writer, st[len(st) - 1]) 	//st.top()
	}
	
	st = st[:len(st)-1] 						//st.pop()
	fmt.Fprintln(writer, st)
}
