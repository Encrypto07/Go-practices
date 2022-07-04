package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	defer wg.Wait()
	finishC := make(chan bool)
	wg.Add(1)
	go func() {
		select {
		case <-time.After(1 * time.Hour):
		case <-finishC:
		}
		wg.Done()
	}()

	t0 := time.Now()
	finishC <- true
	fmt.Printf("waited %v for goroutine to stop\n", time.Since(t0))
}
