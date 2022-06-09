package main

import (
	"fmt"
)

func main() {
	m := make(map[string]string)
	m["fname"] = "sultan"
	m["lname"] = "sheikh"
	
	delete(m, "fname")

	var val, ok = m["fname"]
	var val1, ok1 = m["person"]
	fmt.Println(val, ok)
	fmt.Println(val1, ok1)


}
