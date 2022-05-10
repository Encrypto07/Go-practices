package main

import "fmt"

func main() {
	foo()

}

//No Parameter no return type
func foo() {
	fmt.Println("Hello sultan")
	bar("Manisha")

	x, y := woo("Ashutosh", "singh")
	fmt.Println(x, y)
}

func bar(s string) {
	fmt.Println("Hii", "", s)
}

func woo(fn, ln string) (string, bool) {
	a := fmt.Sprint(fn, " ", ln, " ",`Say Hello`)
	b := true
	return a, b
}
