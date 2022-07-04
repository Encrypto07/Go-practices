package main

//this program illustrate us how we can send the struct with channel

import (
	"fmt"
	"sync"
)

type person struct {
	name string
	age  int
}

func main() {
	var wg sync.WaitGroup
	defer wg.Wait()
	ch := make(chan person)

	wg.Add(2)
	go send(ch, &wg)
	go received(ch, &wg)
}

func send(c chan<- person, wg *sync.WaitGroup) {
	defer wg.Done()
	c <- person{name: "Shubham", age: 19}
	close(c)
}

func received(c <-chan person, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(<-c)
}
