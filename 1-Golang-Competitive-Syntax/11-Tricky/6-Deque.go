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

	//-----------DEQUE----------------
	var dq []int  								//deque<int> dq
	
	for i := 0; i < n; i++ {
		dq = append(dq, arr[i]) 				//dq.push_back(arr[i])
	}
	dq = append([]int{999}, dq...)				//dq.push_front(999)
	fmt.Fprintln(writer, dq)

	if len(dq) > 0 {							//dq.empty()
		fmt.Fprintln(writer, dq[len(dq) - 1]) 	//dq.back()
		fmt.Fprintln(writer, dq[0]) 			//dq.front()
	}

	dq = dq[1:]									//dq.pop_front()
	dq = dq[:len(dq)-1] 						//dq.pop_back()
	fmt.Fprintln(writer, dq)
}
