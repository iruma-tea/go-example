package main

import (
	"fmt"
	"reflect"
	"runtime"
	"unsafe"
)

func main() {
	s := "hello"
	aHdr := (*reflect.StringHeader)(unsafe.Pointer(&s))
	fmt.Println(aHdr.Len)
	for i := 0; i < aHdr.Len; i++ {
		bp := *(*byte)(unsafe.Pointer(aHdr.Data + uintptr(i)))
		fmt.Print(string(bp))
	}
	fmt.Println()
	runtime.KeepAlive(s)
}
