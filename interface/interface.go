package main

import "fmt"

type college_people interface {
	setinfo()
	geinfo()

}

type Teacher struct {
	id     int
	f_name string
	l_name string
}

type student struct {
	id      int
	f_name  string
	l_name  string
	roll_no int
}

func (t *Teacher) setinfo() {
	fmt.Scan(&t.id, &t.f_name, &t.l_name)
}

func (t *Teacher) geinfo() {
	fmt.Println(t.id, t.f_name, t.l_name)
}
func (s *student) setinfo() {
	fmt.Scan(&s.id, &s.f_name, &s.l_name, &s.roll_no)
}
func (s *student) geinfo() {
	fmt.Println(s.id, s.f_name, s.f_name, s.roll_no)
}

func main() {
	t := Teacher{}
	s := student{}

	p := []college_people{&t, &s}
	for _, cp := range p {
		cp.setinfo()
		cp.geinfo()
	}

}
