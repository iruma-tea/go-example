package main

import (
	"fmt"
	"reflect"
)

type Foo struct {
	A int
	B string
}

var x Foo

func DoSomething(f Foo) {
	fmt.Println(f.A, f.B)
}

func main() {
	fmt.Println("===== 14.1 リフレクションを使った実行時の型の処理 =====")

	fmt.Println("----- TypeOf -----")
	{
		v := 0
		w := 0.1

		vType := reflect.TypeOf(v)
		fmt.Println(vType)
		wType := reflect.TypeOf(w)
		fmt.Println(wType)
	}

	fmt.Println("----- ValueOf -----")
	{
		v := 1.03
		vValue := reflect.ValueOf(v)
		fmt.Println(vValue)
	}

	{
		s := []string{"a", "b", "c"}
		sv := reflect.ValueOf(s)
		s2 := sv.Interface().([]string)
		fmt.Println(sv)
		fmt.Printf("%T\n", sv)
		fmt.Println(s2)
		fmt.Printf("%T\n", s2)
	}

	fmt.Println("----- 値 -----")
	{
		v := 1.03
		vValue := reflect.ValueOf(v)
		fmt.Println("vValue:", vValue)
	}

	{
		i := 10
		iv := reflect.ValueOf(&i)
		ivv := iv.Elem()

		ivv.SetInt(20)
		fmt.Println("i:", i)
	}
}
