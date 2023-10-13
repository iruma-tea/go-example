package main

import "fmt"

type Integer interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

func plusOneThousand[T Integer](in T) T {
	// return in + 1_000 // int8に収まらないのでエラー
	return in
}

func main() {
	var a int = 10
	b := plusOneThousand(a)
	fmt.Printf("%T: %v\n", b, b)
}
