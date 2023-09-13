package main

import "fmt"

func main() {
	eventVals := []int{2, 4, 6, 8, 10, 12}
	for i, v := range eventVals {
		fmt.Println(i, v)
	}
}
