package main

import "fmt"

type Stack[T any] struct { // 型Stackの宣言。ジェネリクス版 //liststart1
	vals []T
}

func (s *Stack[T]) Push(val T) {
	s.vals = append(s.vals, val)
}

func (s *Stack[T]) Pop() (T, bool) {
	if len(s.vals) == 0 {
		var zero T // 変数の宣言。zeroはその型のゼロ値になる
		return zero, false
	}
	top := s.vals[len(s.vals)-1]
	s.vals = s.vals[:len(s.vals)-1]
	return top, true
}

func main() {
	var intStack Stack[int]
	intStack.Push(10)
	intStack.Push(20)
	intStack.Push(30)
	v, ok := intStack.Pop()
	fmt.Println(v, ok) // 30 true
	// intStack.Push("nope") //liststart2
	// listend2
}
