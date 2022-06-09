package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	p, err := os.Create("new.txt")
	if err != nil {
		log.Println("Unable to open the file")
	}
	defer p.Close()
	var input string
	fmt.Print("enter input to write to file: ")
	data, _ := fmt.Scan(&input)
	x := strconv.Itoa(data)

	fmt.Println(x)
	// r := strings.NewReader()
	// io.Copy(p, r)

}
