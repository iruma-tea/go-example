package main

import "fmt"

func main() {
	type person struct {
		FirstName  string
		MiddleName *string
		LastName   string
	}

	p := person{
		FirstName: "Pat",
		// MiddleName: "Perry", // ←コンパイル時のエラー
		LastName: "Peterson",
	}
	fmt.Println(p)
}
