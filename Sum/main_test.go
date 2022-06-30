package main

import "testing"

func TestMySum(t *testing.T) {
	a := mySum(3, 2)
	if a != 5 {
		t.Error("Excepted", 5, "Got", a)
	}
}

func TestSubs(t *testing.T) {
	a := subs(5, 2)
	if a != 3 {
		t.Error("Excepted", 3, "Got", a)
	}
}
