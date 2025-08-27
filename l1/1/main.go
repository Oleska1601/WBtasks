package main

import "fmt"

type Human struct {
	name    string
	surname string
	age     int
}

type Action struct {
	Human
}

func (h Human) getFullName() string {
	return h.name + " " + h.surname
}

func main() {
	act := Action{
		Human: Human{
			name:    "Olesya",
			surname: "Frolova",
			age:     19,
		},
	}
	fmt.Println(act.getFullName())
}
