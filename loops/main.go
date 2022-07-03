package main

import "fmt"

func main() {
	xi := []int{1, 2, 3, 4, 5, 6}
	for i := 0; i < len(xi); i++ {
		fmt.Println(i)
	}

	for v := range xi {
		fmt.Println(v)
	}
}
