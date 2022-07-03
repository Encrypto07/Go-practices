package main

import (
	"fmt"
	"sync"
)

func send(c chan<- string) {
	for i := 0; i <= 9; i++ {
		c <- "Hello"
	}
	close(c)
	wg.Done()
}

func received(c <-chan string) {
	for v := range c {
		fmt.Println(v)
	}
	wg.Done()
}

var wg sync.WaitGroup

func main() {
	ch := make(chan string)
	wg.Add(1)
	for i := 0; i <= 10; i++ {
		go send(ch)
	}

	for i := 0; i <= 10; i++ {
		go received(ch)
	}

	wg.Wait()

}
