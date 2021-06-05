package main

import (
	"fmt"
)

func main() {

	arr := [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr)

	arr2 := [...]int{1, 2, 3, 34, 4, 4, 5, 4, 54, 6, 56, 5, 76, 76}
	fmt.Println(arr2)
	fmt.Println(len(arr2))

	arr3 := [5]int{
		0: 1,
		1: 2,
		3: 3,
		2: 4,
		4: 100 + 100,
	}
	fmt.Println(arr3)

	var arr4 [4]int
	arr4[0] = 12
	arr4[1] = 13
	arr4[2] = 144
	arr4[3] = 454
	fmt.Println(arr4)
	add := *&arr
	fmt.Println(add)

	var a int = 12
	fmt.Println(a)
}
