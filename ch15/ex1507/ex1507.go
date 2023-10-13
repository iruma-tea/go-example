package main

import "fmt"

type ImpossiblePrintableInt interface {
	int
	String() string
}

type ImpossibleStruct[T ImpossiblePrintableInt] struct {
	val T
}

type MyInt int

func (mi MyInt) String() string {
	return fmt.Sprint(mi)
}

func main() {
	// s := ImpossibleStruct[int]{10}    //liststart2 // コンパイルエラー
	// s2 := ImpossibleStruct[MyInt]{10} //listend2　 // コンパイルエラー
	// fmt.Println(s)
	// fmt.Println(s2)
}
