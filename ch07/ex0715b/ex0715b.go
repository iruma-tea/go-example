package main

import "fmt"

type MyInt int

func main() {
	var i interface{}
	var mine MyInt = 20
	i = mine
	i3 := i.(string) // iをstring型だと仮定して値をもらおうとするが... パニック！
	fmt.Println(i3)
}
