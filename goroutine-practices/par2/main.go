package main

import (
	"fmt"
	"unsafe"
)

func main() {

	var arr [10]struct{}
	fmt.Println(arr)
	fmt.Println(unsafe.Sizeof(arr))

	
}
