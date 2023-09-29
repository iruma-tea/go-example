package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x int
	xpt := reflect.TypeOf(&x)
	fmt.Println(xpt.Name())        // ""空文字
	fmt.Println(xpt.Kind())        // ptr
	fmt.Println(xpt.Elem().Name()) // int
	fmt.Println(xpt.Elem().Kind()) // int

	s := xpt.Elem().Name()
	fmt.Println(reflect.TypeOf(s)) // string
	c := xpt.Elem().Kind()
	fmt.Println(reflect.TypeOf(c)) // reflect.kind
}
