package main

import "fmt"

func main() {
	i, j := 1, 2
	fmt.Println(&i, &j)

	p := &i
	fmt.Printf("%T\n", p)
	fmt.Printf("type is %T\nand value is %v\n", p, *p)

	*p = 2
	fmt.Printf("Type of *p is %T\n and value stored in variable i is %d\n", *p, i)
	//fmt.Println(i)
}
