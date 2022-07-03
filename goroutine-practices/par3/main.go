/* if we loop twice on same channel then in first loop its have the data but in second time it dont't have the data*/

package main

import "fmt"

func main() {
	ch := make(chan bool, 2)

	ch <- true
	ch <- true
	close(ch)
	for i := 0; i < cap(ch)+1; i++ {
		v, ok := <-ch
		fmt.Println(v, ok)
	}

	fmt.Println("********************************************")
	for j := 0; j < cap(ch); j++ {
		v, ok := <-ch
		fmt.Println(v, ok)
	}

}
