package main

import "fmt"

func main() {

	//Additon of two array elements
	arr1 := [5]int{1, 2, 3, 4, 5}
	fmt.Println("The First Array is =", arr1)

	arr2 := [5]int{6, 7, 8, 9, 10}
	fmt.Println("The second Array is =", arr2)

	//Mathod 1
	arr3 := [5]int{
		0: arr1[0] + arr2[0],
		1: arr1[1] + arr2[1],
		2: arr1[2] + arr2[2],
		3: arr1[3] + arr2[3],
		4: arr1[4] + arr2[4],
	}
	fmt.Println("The First Method Of Printing Array is = ", arr3)

	//Second Method
	var arr4 [5]int
	arr4[0] = arr1[0] + arr2[0]
	arr4[1] = arr1[1] + arr2[1]
	arr4[2] = arr1[2] + arr2[2]
	arr4[3] = arr1[3] + arr2[3]
	arr4[4] = arr1[4] + arr2[4]
	fmt.Println("The Second Method of Printing Array is =", arr4)
}
