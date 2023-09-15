package main

import "fmt"

func main() {

	type person struct {
		FirstName  string
		MiddleName *string
		LastName   string
	}

	p := person{
		FirstName:  "Pat",
		MiddleName: stringp("Perry"),
		LastName:   "Peterson",
	}

	fmt.Println(p)
	fmt.Println(*p.MiddleName)

}

func stringp(s string) *string {
	return &s
}
