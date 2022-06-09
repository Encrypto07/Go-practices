package main

import "fmt"

func main() {
	var fname, sname, address string
	// fmt.Print("Enter your first name: ")
	// fmt.Scan(&fname)
	// fmt.Print("Enter your second name: ")
	// fmt.Scan(&sname)
	// fmt.Print("Enter your address: ")
	// fmt.Scan(&address)

	fmt.Print("Enter the details: ")
	fmt.Scan(&fname, &sname, &address)
	fmt.Printf("The person first name is %s and second name is %s and his/her address is %s\n", fname, sname, address)
}
