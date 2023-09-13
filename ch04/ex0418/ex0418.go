package main

import "fmt"

func main() {
	eventVals := []int{2, 4, 6, 8, 10, 12}
	for _, v := range eventVals {
		v *= 2
	}
	fmt.Println(eventVals)
}
