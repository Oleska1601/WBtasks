package main

import "fmt"

type human struct {
	name    string
	surname string
	age     int
}

type action struct {
	human
}

func (h human) getFullName() string {
	return h.name + " " + h.surname
}

func main() {
	act := action{
		human: human{
			name:    "Olesya",
			surname: "Frolova",
			age:     19,
		},
	}
	fmt.Println(act.getFullName())
}
