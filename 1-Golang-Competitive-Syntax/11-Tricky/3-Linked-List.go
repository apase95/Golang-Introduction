package main

import (
	"bufio"
	"fmt"
	"os"
)

type Node struct {
	val int
	next *Node
}

func pushFront(head *Node, x int) *Node {
	newNode := &Node{
		val: x,
		next: head,
	}
	return newNode
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

	var arr = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &arr[i])
	}

	//-----------LINKED LIST----------------
	var head *Node = nil
	for i := 0; i < n; i++ {
		head = pushFront(head, arr[i])
	}

	for curr := head; curr != nil; curr = curr.next {
		fmt.Fprintf(writer, "%d -> ", curr.val)
	}
	fmt.Fprintln(writer, "NULL")
}