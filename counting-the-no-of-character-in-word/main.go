package main

import (
	"fmt"
)

func main() {
	var word string = "Hello"
	a, _ := fmt.Println(word)

	//Calculating the No of character in the words.
	//Explatation as we know fmt.Println() return two value 1 is number of bytes which is type int and second is error
	//wher n types of int represent an arguments and number of character added thats why we substract the 1 from a
	c := a - 1

	fmt.Println("The number of character in the above word is ", c)
}
