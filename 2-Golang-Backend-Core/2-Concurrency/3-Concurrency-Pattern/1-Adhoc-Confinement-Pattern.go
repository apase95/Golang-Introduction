package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	data := make([]int, 0)

	appendData := func(wg *sync.WaitGroup, data *[]int, value int) {
		defer wg.Done()
		*data = append(*data, value)
	}

	wg.Add(1)
	go appendData(&wg, &data, 10)

	wg.Add(1)
	go appendData(&wg, &data, 20)
	
	wg.Wait()
	fmt.Print(data)
}