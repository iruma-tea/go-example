package main

import "fmt"

type Score int
type HighScore Score

type Person struct {
	LastName  string
	FirstName string
	Age       int
}

type Employee Person

func (s Score) Double() Score {
	return s * 2
}

func main() {
	var i int = 300
	var s Score = 100
	var hs HighScore = 200
	// hs = s // コンパイルエラー
	// s = i // コンパイルエラー
	s = Score(i)
	hs = HighScore(s)
	fmt.Println(s, hs)
	hhs := hs + 20
	fmt.Println(hhs)

	s = 200
	hs = 300
	fmt.Println(s.Double())
	fmt.Println(Score(hs).Double())
	// fmt.Println(hs.Double()) // コンパイルエラー
}
