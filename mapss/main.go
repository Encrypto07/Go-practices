package main

import "fmt"

func main() {
	mx := map[string]string{
		"fname":   "sultan",
		"lname":   "sheikh",
		"address": "hayat nagar",
	}
	fmt.Printf("The first type is %T\n","The second type is %T\n",mx["abc"])

	_, val := mx["person"]

	if val {
		fmt.Println(val)
	} else {
		fmt.Println(val)
	}
}
