package main

import "fmt"

func main() {
	fmt.Println(mySum(2, 3))
	fmt.Println(subs(5, 2))

}

func mySum(num ...int) int {
	sum := 0
	for _, v := range num {
		sum += v
	}
	return sum
}

func subs(num1, num2 int) int {
	c := num1 - num2
	return c
}
