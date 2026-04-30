package main

import (
	"bufio"
	"fmt"
	"os"
)

type DSU struct {
	parent []int
	size []int
}

func NewDSU(n int) *DSU {
	d := &DSU{
		parent: make([]int, n + 1),
		size: make([]int, n + 1),
	}
	for i := 1; i <= n; i++ {
		d.parent[i] = i
		d.size[i] = 1
	}
	return d
}

func (d *DSU) Find(u int) int {
	if u == d.parent[u] {
		return u
	}
	d.parent[u] = d.Find(d.parent[u])
	return d.parent[u]
}

func (d *DSU) Union(u, v int) bool {
	rootU := d.Find(u)
	rootV := d.Find(v)
	if rootU != rootV {
		if d.size[rootU] < d.size[rootV] {
			rootU, rootV = rootV, rootU
		}
		d.parent[rootV] = rootU
		d.size[rootU] += d.size[rootV]
		return true
	}
	return false
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

	var n, m int
	fmt.Fscan(reader, &n, &m)
	
	dsu := NewDSU(n)
	for i := 0; i < m; i++ {
		var u, v int
		fmt.Fscan(reader, &u, &v)
		if dsu.Union(u, v) {
			fmt.Fprintf(writer, "[%d -> %d]\n", u, v)
		} else {
			fmt.Fprintf(writer, "[%d -> %d]: Already Connected\n", u, v)
		}
	}
}