package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	eve := make(chan int)
	odd := make(chan int)
	quid := make(chan int)
	wg.Add(2)
	go send(eve, odd, quid, &wg)
	go received(eve, odd, quid, &wg)
	wg.Wait()
}

func send(e, o, q chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 100; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	q <- 0
	close(e)
	close(o)
}

func received(e, o, q <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case v := <-e:
			fmt.Println("This is even number", v)
		case v := <-o:
			fmt.Println("This is odd number", v)
		case v := <-q:
			fmt.Println("This is quit", v)
			return
		}
	}
}
