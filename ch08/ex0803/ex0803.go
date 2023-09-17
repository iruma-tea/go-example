package main

import (
	"fmt"
	"os"
)

func doubleEvent(i int) (int, error) {
	if i%2 != 0 {
		return 0, fmt.Errorf("%dは偶数ではありません", i)
	}

	return i * 2, nil
}

func main() {
	i := 19
	double, err := doubleEvent(i)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(double)
}
