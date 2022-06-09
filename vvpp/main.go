package main

import (
	"fmt"
)

func main() {
	p := sum(1, 0)
	fmt.Println(p)

	l := rev("ABC")
	fmt.Println(l)
	fmt.Println(len("SSS"))
	a := sl([]int{1, 2, 3})
	fmt.Println(a)

	name, val := hoo("sultan", 12, 12)
	fmt.Printf("The name and value is %s\t%d", name, val)
}

func sum(num ...int) int {
	temp := 0
	for _, v := range num {
		temp += v
	}
	return temp
}

func rev(s string) (res string) {
	for _, v := range s {
		res = string(v) + res
	}
	return
}

func sl(s []int) []int {
	// s = []int{1, 2, 3, 4}
	return s

}

func hoo(s string, num ...int) (string, int) {
	s = string(rev(s))
	temp := 0
	for _, v := range num {
		temp += v
	}
	return s, temp
}

