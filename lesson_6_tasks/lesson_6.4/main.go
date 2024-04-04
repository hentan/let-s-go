package main

import "fmt"

type contacts struct {
	Addss string
	Phone string
}

type user struct {
	ID       int
	Name     string
	contacts
}

type employee struct {
	ID       int
	Name     string
	contacts
}

func main() {
	u := user{
		ID:   1,
		Name: "Vasya",
		contacts: contacts{
			Addss: "123 st",
			Phone: "12345",
		},
	}

	e := employee{
		ID:   2,
		Name: "Petya",
		contacts: contacts{
			Addss: "223 st",
			Phone: "22345",
		},
	}

	fmt.Println(u.Addss, u.Phone, e.Addss, e.Phone)
}