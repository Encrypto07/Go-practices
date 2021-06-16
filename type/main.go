package main

import "fmt"

// we can use `` to for high level string formating but
var myname string = `Sultan   Is a good buy 



and he is also a good person`

var myage int32 = 24
var myheight float32 = 173.32

func main() {

	fmt.Println("My name is a", myname)

	//we can use double \n to do more next line   and %T is used to check the type of
	fmt.Printf("Its Type is %T\n\n", myname)

	fmt.Println("My age is a", myage)
	fmt.Printf("Its Type is %T\n", myage)

	fmt.Println("My height is a", myheight)
	fmt.Printf("Its Type is %T\n", myheight)

	fmt.Print(len(myname))

}
