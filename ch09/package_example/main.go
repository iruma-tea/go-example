package main

import (
	"fmt" // パッケージprint

	print "example.co.jp/package_example/fomatter" // パッケージmath
	"example.co.jp/package_example/math"           // パッケージmath
)

func main() {
	num := math.Double(2)       // パッケージmath（math/math.go）
	output := print.Format(num) // パッケージprint（formatter/formatter.go）
	fmt.Println(output)
}
