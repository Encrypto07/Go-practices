package main

import "fmt"

//Creating Our Own Type
type HotDog int

//Declearing and assigning a Age Variable with Type HotDog
var Age HotDog = 33

func main() {
	fmt.Println(Age)
	fmt.Printf("%T", Age)
}
