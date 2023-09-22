package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	funcs := prepareFunctions()
	s := "Time flies like an arrow."
	r := searchData(s, funcs)
	fmt.Println("結果:", r)

	time.Sleep(1 * time.Second)
	fmt.Println("mainを終了")
}

func searchData(s string, searchers []func(string) []string) []string {
	done := make(chan struct{})
	resultChan := make(chan []string)

	for _, searcher := range searchers {
		go func(f func(string) []string) {
			select {
			case resultChan <- f(s):
				fmt.Println("結果が戻ってきた")
			case <-done:
				fmt.Println("doneを選択")
			}
		}(searcher)
	}
	r := <-resultChan
	close(done)
	return r
}

func prepareFunctions() []func(string) []string {
	search1 := func(a string) []string {
		b := strings.ToLower(a)
		fmt.Println("1:", b)
		r := strings.Split(b, " ")
		return r
	}

	search2 := func(a string) []string {
		b := strings.ToUpper(a)
		fmt.Println("2:", b)
		r := strings.Split(b, " ")
		return r
	}

	search3 := func(a string) []string {
		b := strings.ReplaceAll(a, "i", "I")
		fmt.Println("3:", b)
		r := strings.Split(b, " ")
		return r
	}

	search4 := func(a string) []string {
		b := strings.ReplaceAll(a, "e", "E")
		a = strings.ReplaceAll(a, "r", "L")
		fmt.Println("4:", b)
		r := strings.Split(b, " ")
		return r
	}

	funcs := []func(string) []string{search1, search2, search3, search4}
	return funcs
}
