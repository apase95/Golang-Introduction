package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

type PQItem struct {
	node int
	dist int
}

type PQ []PQItem

func (pq PQ) Len() int {
	return len(pq)
}

func (pq PQ) Less(i, j int) bool {
	return pq[i].dist < pq[j].dist
}

func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PQ) Push(x any) {
	*pq = append(*pq, x.(PQItem))
}

func (pq *PQ) Pop() any {
	old := *pq
	n := len(old)
	item := old[n - 1]
	*pq = old[0 : n - 1]
	return item
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

	//reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	pq := make(PQ, 0)
	heap.Init(&pq)
	heap.Push(&pq, PQItem{node: 1, dist: 10})
	heap.Push(&pq, PQItem{node: 2, dist: 5})
	heap.Push(&pq, PQItem{node: 3, dist: 15})

	for pq.Len() > 0 {
		top := heap.Pop(&pq).(PQItem)
		fmt.Printf("Node: %d, Distance: %d\n", top.node, top.dist)
	}
}