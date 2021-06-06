package main

import (
	"fmt"
	"log"
)

func main() {
	var name string = "sultan"
	a, err := fmt.Println(name)

	//we can check the conditions explain in the official docs of golang in fmt package in the function of Prinln() it return two values
	//1 an n type of int and second is err type of error
	//if you want to check that Println return value err satisfies that it return an err or not use == instead of !=
	if err != nil {
		log.Fatal("Not able to Print on the console")
	}
	fmt.Println(a, err)
}
