package main

import "fmt"

func main() {
	defer fmt.Println("world")
	defer fmt.Println("one")
	defer fmt.Println("two")
	fmt.Println("hello")
	reverse()
}

func reverse() {

	for i := 1; i <= 9; i++ {
		defer fmt.Println(i)
	}
}
