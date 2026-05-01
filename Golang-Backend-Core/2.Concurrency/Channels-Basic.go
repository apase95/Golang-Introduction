package main

import (
	"fmt"
	"sync"
	"time"
)

func generateLogs(serverName string, logChan chan string, wg *sync.WaitGroup){
	defer wg.Done()
	for i := 1; i <= 3; i++ {
		msg := fmt.Sprintf("Log %d from %s", i, serverName)
		logChan <- msg
		time.Sleep(400 * time.Millisecond)
	}
}

func main() {
	fmt.Println("1. BUFFERED CHANNELS")
	msgChan := make(chan string, 3)
	
	msgChan <- "Message 1"
	msgChan <- "Message 2"
	msgChan <- "Message 3"
	fmt.Println("Pop:", <-msgChan)
	fmt.Println("Pop:", <-msgChan)
	fmt.Println("Pop:", <-msgChan)
	
	fmt.Println("2. MULTIPLE SENDER, ONE RECEIVER")
	logQueue := make(chan string)
	var wg sync.WaitGroup
	
	//2.1
	servers :=[]string{"Web-01", "DB-01"}
	for _,srv := range servers {
		wg.Add(1)
		go generateLogs(srv, logQueue, &wg)
	}
	//2.2
	go func() {
		wg.Wait()
		close(logQueue)
		fmt.Println("[System] All Server was closed")
	}()
	//2.3
	for logMsg := range logQueue {
		fmt.Printf("[Log Collector]: %s\n", logMsg)
	}
}