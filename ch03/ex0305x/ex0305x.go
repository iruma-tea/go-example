package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4} //liststart
	y := x[:2]
	fmt.Println(y) // [1 2]  （xの先頭から2つ。つまりx[0]とx[1]）
	z := x[1:]
	fmt.Println(z) // [2 3 4] （xの2つめから残り全部。x[1]からx[3]まで）
	x[1] = 20
	fmt.Println("x:", x) // x: [1 20 3 4]
	fmt.Println("y:", y) // y: [1 20]
	fmt.Println("z:", z) // z: [20 3 4]
	y[0] = 10            // x[1]も変わる
	fmt.Println("----")
	fmt.Println("x:", x) // x: [10 20 3 4]
	fmt.Println("y:", y) // y: [10 20]    （xの先頭から2つ）
	fmt.Println("z:", z) // z: [20 3 4]   （xの2つめから残り全部）
	z[1] = 30            // x[2]もかわる
	fmt.Println("----")
	fmt.Println("x:", x) // x: [10 20 30 4]
	fmt.Println("y:", y) // y: [10 20]
	fmt.Println("z:", z) // z: [20 30 4]
	// listend
}
