package main

import "fmt"

func main() {
	x := 10
	fmt.Println(x)
	// ↓エラーとなる
	// fmt := "おっと～"
	// fmt.Println(fmt)
}
