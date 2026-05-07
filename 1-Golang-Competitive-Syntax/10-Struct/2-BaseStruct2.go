package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Person struct {
	name string
	age int
	job string
	salary int
}

func comparePeople(p1, p2 Person) bool {
	if p1.salary == p2.salary {
		return p1.age < p2.age
	}
	return p1.salary < p2.salary
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
	markPeople := make([]Person, n)
	
	for i := 0; i < n; i++ {
		var p Person
		fmt.Fscanln(reader, &p.name, &p.age, &p.job, &p.salary)
		markPeople[i] = p
	}

	sort.Slice(markPeople, func(i, j int) bool {
		return comparePeople(markPeople[i], markPeople[j])
	})

	for _, p := range markPeople {
		fmt.Fprintln(writer, p.name, p.age, p.job, p.salary)
	}
}