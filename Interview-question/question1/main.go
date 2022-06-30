package main

import (
	"fmt"
	"sync"
	"time"
)

/*
Interview Question

create 2 groutine the first groutine print 10 message at time duration of 2 second after second goroutine print 10 message of
time duration of 2 second

*/

var wg sync.WaitGroup

func main() {

	ch1 := make(chan string)

	wg.Add(2)
	go print(ch1)
	for v := range ch1 {
		fmt.Println(v)
	}

	ch2 := make(chan string)
	go print2(ch2)
	for v := range ch2 {
		fmt.Println(v)
	}
	wg.Wait()
}

func print(c chan<- string) {
	for i := 0; i <= 10; i++ {
		time.Sleep(1 * time.Second)
		c <- "Hello"

	}
	close(c)
	wg.Done()
}

func print2(c chan<- string) {
	for i := 0; i <= 10; i++ {
		time.Sleep(1 * time.Second)
		c <- "World !"
	}
	close(c)
	wg.Done()
}
