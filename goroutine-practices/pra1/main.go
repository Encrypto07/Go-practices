package main

func main() {
	ch := make(chan struct{})

	go send(ch)
}

func send(ch chan struct{}) {
	p := struct {
		FirstName string
		LastName  string
	}{
		FirstName: "sultan",
		LastName:  "sheikh",
	}
	ch <- struct{}{}
}
