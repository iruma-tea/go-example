package main

import (
	"fmt"
	"reflect"
)

func main() {
	i := 10
	iv := reflect.ValueOf(&i)
	ivv := iv.Elem()
	fmt.Println(ivv)
	ivv.SetInt(20)
	fmt.Println(ivv)
}
