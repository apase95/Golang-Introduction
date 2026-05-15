package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println(i)
		}
	}()

	time.Sleep(time.Second)
}