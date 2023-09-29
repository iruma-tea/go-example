package main

import (
	"fmt"
	"reflect"
)

type Foo struct {
	A int
	B string
}

func DoSomething(f Foo) {
	fmt.Println(f.A, f.B)
}

func main() {
	var x int
	xt := reflect.TypeOf(x)                  // xtはreflect.Type型。xの型に関する情報をもつ
	fmt.Printf("xt.Name():%v|\n", xt.Name()) // int
	f := Foo{}
	ft := reflect.TypeOf(f)
	fmt.Printf("ft.Name():%v|\n", ft.Name()) // Foo
	xpt := reflect.TypeOf(&x)
	fmt.Printf("xpt.Name():%v|\n", xpt.Name()) // "" （何も表示されない）

	fmt.Println(xt.Kind())  // int //liststart3
	fmt.Println(ft.Kind())  // struct
	fmt.Println(xpt.Kind()) // ptr //listend3
}
