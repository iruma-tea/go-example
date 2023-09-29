package main

import (
	"fmt"
	"reflect"
)

func main() {
	f := func(a, b int) int {
		return a + b
	}

	fmt.Println(f(3, 9))

	ft := reflect.TypeOf(f)
	fmt.Println(ft)         // func(int, int) int
	fmt.Println(ft.NumIn()) // 2 // fの引数は2個

	x := 3
	xt := reflect.TypeOf(x)
	fmt.Println(xt)         // int
	fmt.Printf("%T\n", xt)  // *reflect.rtype
	fmt.Println(xt.NumIn()) // パニックになる（xは関数ではない）
}
