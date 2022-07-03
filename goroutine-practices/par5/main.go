package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	defer wg.Wait()
	ch1 := make(chan string)
	ch2 := make(chan string)

	wg.Add(2)
	go print1(ch1, &wg)
	for v := range ch1 {
		fmt.Println(v)
	}

	go print2(ch2, &wg)
	for v := range ch2 {
		fmt.Println(v)
	}

}

func print1(c chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i <= 10; i++ {
		time.Sleep(2 * time.Second)
		c <- "Hello"
	}
	close(c)
}

func print2(c chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i <= 10; i++ {
		time.Sleep(2 * time.Second)
		c <- "World !"
	}
	close(c)
}
