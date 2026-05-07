package main

import (
	"bufio"
	"fmt"
	"os"
)

type User struct {
	Name string
	Age  int
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

	var Name string
	var Age int
	fmt.Fscan(reader, &Name, &Age)

	var user = new(User)
	user.Name = Name
	user.Age = Age

	fmt.Fprint(writer, user)
}