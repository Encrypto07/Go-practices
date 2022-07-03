package main

import "fmt"

func checkmonth(n int)  {
	month := map[int]string{
		1: "monday",
		2: "tuesday",
		3: "wednesday",
		4: "thrusday",
		5: "fridaay",
		6: "saturday",
		7: "sunday",
	}
	if v, k := month[n]; k {
		fmt.Println(v)
	}else {
		fmt.Println("key is not presend in map",)
	}
}

func main() {
	checkmonth(7)
	
}
