package main

import "fmt"

func main() {
	var str = "Hello"

	fmt.Println("Before the reversing a string is ", str)
	var result string = ""

	for i := len(str) - 1; i >= 0; i-- {

		result = result + string(str[i])

	}
	fmt.Println("After reversing the string is ", result)
}
