package main

import "fmt"

func main() {
	var i any
	i = 20
	fmt.Println(i)
	i = "Hello"
	fmt.Println(i)
	i = struct {
		FirstName string
		LastName  string
	}{
		"信玄",
		"武田",
	}
	fmt.Println(i)

	var i2 any
	i2 = 20
	fmt.Println(i2)
	i2 = "Hello"
	fmt.Println(i2)
	i2 = struct {
		FirstName string
		LastName  string
	}{"信玄", "武田"}
	fmt.Println(i2)

	type Person struct {
		FirstName string
		LastName  string
	}

	var i3 any
	i3 = 20
	fmt.Println(i3)
	i3 = "Hello"
	fmt.Println(i3)
	i3 = Person{
		FirstName: "信玄",
		LastName:  "武田",
	}
	i4 := i3.(Person)
	fmt.Println(i4.LastName, i4.FirstName)
}
