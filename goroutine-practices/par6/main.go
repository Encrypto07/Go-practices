package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	defer wg.Wait()
	ch := make(chan string)

	wg.Add(2)
	go send(ch, &wg)
	go received(ch, &wg)
}

func send(c chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 10; i++ {
		time.Sleep(3 * time.Second)
		c <- fmt.Sprintf("Message %d", i)
	}
	close(c)
}

func received(c <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for v := range c {
		fmt.Println(v)
	}
}
