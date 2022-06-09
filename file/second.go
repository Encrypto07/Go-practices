package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	f, err := os.Open("new.txt")
	if err != nil {
		log.Printf("Unable to open a file %v", err)
		return
	}
	defer f.Close()

	bs, err := ioutil.ReadAll(f)
	if err != nil {
		log.Printf("unable to read %v\n", err)
		return
	}
	fmt.Println(string(bs))

}
