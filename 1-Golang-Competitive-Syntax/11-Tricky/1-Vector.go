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

	//-----------VECTOR----------------
	var v []int  					//vector<int> v
	
	for i := 0; i < n; i++ {
		v = append(v, arr[i]) 		//v.push_back(arr[i])
	}
	fmt.Fprintln(writer, v)

	v = v[:len(v)-1] 				//v.pop_back()
	fmt.Fprintln(writer, v)

	clear(v) 						//v.clear()
	fmt.Fprintln(writer, v)
}