package main

import "fmt"

type Stack[T any] struct {
	vals []T
}

func (s *Stack[T]) Push(val T) {
	s.vals = append(s.vals, val)
}

func (s *Stack[T]) Pop() (T, bool) {
	if len(s.vals) == 0 {
		var zero T
		return zero, false
	}
	top := s.vals[len(s.vals)-1]
	s.vals = s.vals[:len(s.vals)-1]
	return top, true
}

func (s Stack[T]) Contains(val T) bool { //liststart2
	for _, v := range s.vals {
		// any 型の値を保存し、それを取り出すことしかできません。 ==を使うには別の型が必要です。ほとんどすべての Go
		//の型は==と !=で比較可能（comparable）なので、 comparable という新しい組み込みのインタ
		//フェースがユニバースブロックで定義されています。
		// if v == val {
		println(v)
		return true
		// }
	}
	return false
} //listend2

func main() {
	var intStack Stack[int]
	intStack.Push(10)
	intStack.Push(20)
	intStack.Push(30)
	fmt.Println(intStack.Contains(10)) // true
	fmt.Println(intStack.Contains(5))  // false
}
