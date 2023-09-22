package main

import (
	"fmt"
	"time"
)

func countTo(max int) <-chan int {
	ch := make(chan int)
	go func() {
		for i := 0; i < max; i++ {
			ch <- i
		}
		close(ch)
	}()
	return ch
}

func main() {
	for i := range countTo(10) {
		fmt.Print(i, " ")
	}
	fmt.Println()
	doSomethingTakingLogTime() // 時間のかかる処理。無名関数の実行は終わっている
}

func doSomethingTakingLogTime() {
	time.Sleep(5 * time.Second)
}
