package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	defer wg.Wait()
	c1 := make(chan string)

	for i := 0; i <= 4; i++ {
		wg.Add(1)
		go send(c1, &wg)
	}
	for j := 0; j <= 4; j++ {
		go received(c1, &wg)
	}

}

func send(c chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i <= 10; i++ {
		c <- fmt.Sprintf("MEssage %d", i)
	}
}

func received(c <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for v := range c {
		fmt.Println(v)
	}
}
