package main

import "fmt"

func main() {
	eventVals := []int{2, 4, 6, 8, 10}
	for i := 1; i < len(eventVals)-1; i++ {
		fmt.Println(i, eventVals[i])
	}
}
