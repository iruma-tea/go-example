package main

import (
	"fmt"
	"reflect"
)

func main() {
	// 1. まず reflect.TypeOf()の値を得る（reflection的型情報を得る）
	// 2. その型に定義されているメソッドを使う（Name、Kindなど）

	type Foo struct {
		A int
		B string
	}

	var x int
	xt := reflect.TypeOf(x)
	fmt.Println(xt.Name())
	fmt.Printf("%T\n", xt.Name())

	f := Foo{}
	ft := reflect.TypeOf(f)
	fmt.Println(ft.Name())
	fmt.Printf("%T\n", ft.Name())

	xpt := reflect.TypeOf(&x)
	fmt.Println(xpt.Name())
	fmt.Printf("%T\n", xpt.Name())

	fmt.Println("-----")
	fmt.Println(xt.Kind())
	fmt.Println(ft.Kind())
	fmt.Println(xpt.Kind())

	fmt.Println("-----")
	fmt.Printf("%T（%d）\n", xt.Kind(), xt.Kind())   // reflect.Kind(2) // iotaを使って定義された定数 具体的な値は2
	fmt.Printf("%T(%d)\n", ft.Kind(), ft.Kind())   // reflect.Kind(25)
	fmt.Printf("%T(%d)\n", xpt.Kind(), xpt.Kind()) // reflect.Kind(22)

}
