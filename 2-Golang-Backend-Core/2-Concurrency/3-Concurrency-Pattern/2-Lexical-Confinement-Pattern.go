package main

import (
	"fmt"
)

func lexicalConfinement() <-chan int {
	results := make(chan int)

	go func() {
		defer close(results)
		data := []int{10, 20, 30}

		for _, val := range data {
			results <- val 
		}
	}()

	return results
}

func main() {
	resultsChan := lexicalConfinement()

	var finalData []int
	for val := range resultsChan {
		finalData = append(finalData, val)
	}

	fmt.Println(finalData)
}